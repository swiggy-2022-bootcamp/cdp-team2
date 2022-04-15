package services

type ICheckoutService interface {
	StartCheckout(cartId int) error
	ApplyReward(orderId, points int) error
	ConfirmOrder(orderId int) error
	GetOrder(orderId int) error
}

type CheckoutService struct {
}

func NewCheckoutService() *CheckoutService {
	// create grpc connections
	return &CheckoutService{}
}

func (cs *CheckoutService) StartCheckout(cartId int) error {
	//Get Cart [Cart MS]

	//Deduct Cart Products from Invertory [Product MS]
	//if any product is not available fail order

	//create order obj with status pending
	//Calculate total

	//Create Order [Order MS]

	//Clear Cart [Cart MS]

	//return order to user
	return nil
}

func (cs *CheckoutService) ApplyReward(orderId, points int) error {

	//check order status = pending [Order MS]

	//apply reward [Reward MS]

	//if final price == 0 change order status to Confirmed [Order MS]

	//return order to user
	return nil
}

func (cs *CheckoutService) ConfirmOrder(orderId int) error {
	//check if shipping address is set

	//change order status to Confirmed [Order MS]

	//return order to user
	return nil
}

func (cs *CheckoutService) GetOrder(orderId int) error {
	// get order from [Order MS]
	return nil
}
