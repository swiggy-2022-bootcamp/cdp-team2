package grpc_services

import (
	"context"
	"errors"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Order/protos/order"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
)

type Server struct {
	*pb.UnimplementedServiceServer
}

func (s *Server) OrderCreateService(ctx context.Context, order *pb.OrderCreateRequest) (*pb.OrderCreateResponse, error) {

	customerId, status, addressId, totalPrice, payedPrice, productDesc := order.Order.CustomerId, order.Order.Status, order.Order.AddressId, order.Order.TotalPrice, order.Order.PayedPrice, order.Order.ProductDesc

	var productDescArray []models.ProductDesc

	for i := 0; i < len(productDesc); i++ {
		//productDescArray = append(productDescArray, productDesc[i])
		productDescArray = append(productDescArray, models.ProductDesc{ProductId: productDesc[i].ProductId, Points: int32(int(productDesc[i].Points)), Reward: int32(int(productDesc[i].Reward)), Quantity: int32(int(productDesc[i].Quantity)), Price: productDesc[i].Price})

	}

	orderCreate := models.Order{
		CustomerId:  customerId,
		Status:      models.Status(status),
		AddressId:   addressId,
		TotalPrice:  totalPrice,
		PayedPrice:  payedPrice,
		ProductDesc: productDescArray,
	}

	orderResponse, err := dao.GetOrderDao().Create(orderCreate)

	if err != nil {
		return nil, errors.New("error occurred")
	}

	order1 := &pb.OrderResponse{
		OrderId:    orderResponse.OrderId,
		CustomerId: orderResponse.CustomerId,
		Status:     int32(orderResponse.Status),
		AddressId:  orderResponse.AddressId,
		TotalPrice: orderResponse.TotalPrice,
		PayedPrice: orderResponse.PayedPrice,
		//ProductDesc: productDescArray,
	}

	return &pb.OrderCreateResponse{
		Order: order1,
	}, nil
}

func (s *Server) OrderStatusUpdateService(ctx context.Context, order *pb.OrderStatusUpdateRequest) (*pb.OrderStatusUpdateResponse, error) {

	orderId, status := order.OrderId, order.Status

	orderUpdate := models.Order{
		Status: models.Status(status),
	}

	_, err := dao.GetOrderDao().UpdateByID(orderId, orderUpdate)

	if err != nil {
		return nil, errors.New("error occurred")
	}

	return &pb.OrderStatusUpdateResponse{
		Response: true,
	}, nil
}

func (s *Server) GetOrderService(ctx context.Context, order *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {

	orderId := order.OrderId

	orderResponse, err := dao.GetOrderDao().GetByID(orderId)

	var productDescArray []*pb.ProductDesc

	for i := 0; i < len(orderResponse.ProductDesc); i++ {
		productDescArray = append(productDescArray, &pb.ProductDesc{ProductId: orderResponse.ProductDesc[i].ProductId, Points: int32(int(orderResponse.ProductDesc[i].Points)), Reward: int32(int(orderResponse.ProductDesc[i].Reward)), Quantity: int32(int(orderResponse.ProductDesc[i].Quantity)), Price: orderResponse.ProductDesc[i].Price})

	}

	order1 := &pb.OrderResponse{
		OrderId:     orderResponse.OrderId,
		CustomerId:  orderResponse.CustomerId,
		Status:      int32(orderResponse.Status),
		AddressId:   orderResponse.AddressId,
		TotalPrice:  orderResponse.TotalPrice,
		PayedPrice:  orderResponse.PayedPrice,
		ProductDesc: productDescArray,
	}

	if err != nil {
		return nil, errors.New("error occurred")
	}

	return &pb.GetOrderResponse{
		Order: order1,
	}, nil
}
