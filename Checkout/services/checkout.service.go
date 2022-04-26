package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	cartpb "common/protos/cart"
	orderpb "common/protos/order"
	productpb "common/protos/products"
	rewardpb "common/protos/reward"

	g "github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/grpc"
)

const (
	ORDER_DECLINED  = iota + 1
	ORDER_PENDING   = iota + 1
	ORDER_CONFIRMED = iota + 1
)

type ICheckoutService interface {
	StartCheckout(customerId string) (*orderpb.OrderResponse, error)
	ApplyReward(orderId string, points int) (*orderpb.OrderResponse, error)
	ConfirmOrder(orderId string) (*orderpb.OrderResponse, error)
	GetOrder(orderId string) (*orderpb.OrderResponse, error)
}

type CheckoutService struct {
	orderClient   *g.OrderClient
	productClient *g.ProductClient
	rewardClient  *g.RewardClient
	cartClient    *g.CartClient
}

func NewCheckoutService() (*CheckoutService, error) {
	bg := context.Background()
	oc, err := g.NewOrderClient(bg)
	if err != nil {
		return nil, err
	}

	pc, err := g.NewProductClient(bg)
	if err != nil {
		return nil, err
	}

	rc, err := g.NewRewardClient(bg)
	if err != nil {
		return nil, err
	}

	cc, err := g.NewCartClient(bg)
	if err != nil {
		return nil, err
	}

	return &CheckoutService{oc, pc, rc, cc}, nil
}

func (cs *CheckoutService) StartCheckout(customerId string) (*orderpb.OrderResponse, error) {
	//Get Cart [Cart MS]

	cartresp, err := cs.cartClient.GetCart(cs.cartClient.CtxWithTimeOut(), &cartpb.CartRequest{CustomerId: customerId})
	if err != nil {
		log.Printf("[Error] Could not fetch cart")
		return nil, err
	}

	log.Printf("[GRPC Call] cartresp %+v", cartresp)

	quantityMap := make(map[int64]int32)

	//Deduct Cart Products from Invertory [Product MS]
	var cartItems []*productpb.ProductsIDAndQnty

	if cartresp.Cart == nil || cartresp.Cart.Products == nil || len(cartresp.Cart.Products) == 0 {
		e := fmt.Errorf("cart is empty")
		log.Printf("[INFO] %s", e.Error())
		return nil, e
	}

	for _, p := range cartresp.Cart.Products {

		pid, err := strconv.Atoi(p.ProductId)
		if err != nil {
			log.Printf("[Error] Invalid Product id in cart : %s", err.Error())
			return nil, err
		}

		cartItems = append(cartItems, &productpb.ProductsIDAndQnty{ProductID: int64(pid), Quantity: uint64(p.Quantity)})
		quantityMap[int64(pid)] = p.Quantity
	}

	res, err := cs.productClient.DeductProductsQuantity(cs.productClient.CtxWithTimeOut(), &productpb.CheckoutRequest{ProductsIDAndQnty: cartItems})
	if err != nil {
		log.Printf("[Error] couldn't adjust inventory = %s", err.Error())
		return nil, err
	}

	log.Printf("[GRPC Call] product reps %+v", res)

	totalPrice := 0.0

	var orderItems []*orderpb.ProductDesc

	if res.AvailableProducts != nil {

		for _, prod := range res.AvailableProducts {
			prodresp, err := cs.productClient.GetProductById(cs.productClient.CtxWithTimeOut(), &productpb.ProductIDRequest{ProductID: prod.ProductID})

			log.Printf("[GRPC Call] product single reps %+v", prodresp)

			if err != nil {
				log.Printf("[Error] couldn't find product = %d", prod.ProductID)
				return nil, err
			}

			p, err := strconv.ParseFloat(prodresp.Product.Price, 32)
			if err != nil {
				log.Printf("[Error] invalid product price = %s", prodresp.Product.Price)
				return nil, err
			}

			orderItems = append(orderItems, &orderpb.ProductDesc{
				ProductId: int32(prodresp.Product.ID),
				Points:    int32(prodresp.Product.Points),
				Reward:    int32(prodresp.Product.Rewards),
				Quantity:  quantityMap[prodresp.Product.ID],
				Price:     float32(p),
			})

			totalPrice += (p * float64(quantityMap[prodresp.Product.ID]))
		}

	}

	//create order obj with status pending
	//Calculate total

	//Create Order [Order MS]
	orderResp, err := cs.orderClient.OrderCreateService(
		cs.orderClient.CtxWithTimeOut(),
		&orderpb.OrderCreateRequest{
			Order: &orderpb.Order{
				CustomerId:  customerId,
				Status:      ORDER_PENDING,
				TotalPrice:  float32(totalPrice),
				ProductDesc: orderItems,
			},
		})

	log.Printf("[GRPC Call] orderResp %+v", orderResp)

	if err != nil {
		log.Printf("[Error] couldn't create order = %s", err.Error())
		return nil, err
	}

	//Clear Cart [Cart MS]
	clearcartres, err := cs.cartClient.EmptyCart(cs.cartClient.CtxWithTimeOut(), &cartpb.CartRequest{CustomerId: customerId})
	if err != nil {
		log.Printf("[Error] Could not clear cart")
		//don't exit because order is already created
	}

	log.Printf("[GRPC Call] clearcartres %+v", clearcartres)

	//return order to user
	return orderResp.Order, nil
}

func (cs *CheckoutService) ApplyReward(orderId string, points int) (*orderpb.OrderResponse, error) {

	//check order status = pending [Order MS]
	iniorder, err := cs.GetOrder(orderId)

	log.Printf("[GRPC Call] iniorder %+v", iniorder)

	if err != nil {
		log.Printf("[Error] invalid order id = %s", orderId)
		return nil, err
	}

	if iniorder.Status != ORDER_PENDING {
		//order is either completed or declined hence can't apply reward
		e := fmt.Errorf("order not in pending state order_id = %s", orderId)
		log.Printf("[Info] %s", e.Error())
		return nil, e
	}

	//apply reward [Reward MS]
	rewardResp, err := cs.rewardClient.RedeemReward(cs.rewardClient.CtxWithTimeOut(), &rewardpb.RedeemRewardRequest{OrderId: iniorder.OrderId, Points: int32(points)})
	log.Printf("[GRPC Call] rewardResp %+v", rewardResp)

	if err != nil {
		log.Printf("[Error] couldn't redeem reward = %s", err.Error())
		return nil, err
	}

	//if final price == 0 change order status to Confirmed [Order MS]
	if rewardResp.FinalPrice == 0.0 {
		// no need to pay
		orderResp, err := cs.orderClient.OrderStatusUpdateService(cs.orderClient.CtxWithTimeOut(),
			&orderpb.OrderStatusUpdateRequest{OrderId: orderId, Status: ORDER_CONFIRMED})

		log.Printf("[GRPC Call] orderResp %+v", orderResp)
		if err != nil || !orderResp.Response {
			log.Printf("[Error] couldn't mark order completed = %s", err.Error())
		}
	}

	//return order to user
	return cs.GetOrder(orderId)
}

func (cs *CheckoutService) ConfirmOrder(orderId string) (*orderpb.OrderResponse, error) {
	//check if shipping address is set
	iniorder, err := cs.GetOrder(orderId)
	if err != nil {
		return nil, err
	}
	log.Printf("[GRPC Call] iniorder %+v", iniorder)

	if iniorder.AddressId == 0 {
		e := fmt.Errorf("Shipping address not set for order = %s", orderId)
		log.Printf("[Info] %s", e.Error())
		return nil, e
	}

	//change order status to Confirmed [Order MS]
	orderResp, err := cs.orderClient.OrderStatusUpdateService(cs.orderClient.CtxWithTimeOut(),
		&orderpb.OrderStatusUpdateRequest{OrderId: orderId, Status: ORDER_CONFIRMED})

	log.Printf("[GRPC Call] orderResp %+v", orderResp)

	if err != nil || !orderResp.Response {
		log.Printf("[Error] couldn't mark order completed = %s", err.Error())
		return nil, err
	}

	//return order to user
	return cs.GetOrder(orderId)
}

func (cs *CheckoutService) GetOrder(orderId string) (*orderpb.OrderResponse, error) {
	// get order from [Order MS]
	orderResp, err := cs.orderClient.GetOrderService(cs.orderClient.CtxWithTimeOut(), &orderpb.GetOrderRequest{OrderId: orderId})

	log.Printf("GRPC Order resp : %+v", orderResp)
	if err != nil {
		log.Printf("[Error] error getting order = %s", orderId)
		return nil, fmt.Errorf("order not found")
	}
	return orderResp.Order, nil
}
