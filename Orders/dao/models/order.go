package models

type ProductDesc struct {
	ProductId int32   `json:"productId,omitempty"  extensions:"x-omitempty"`
	Points    int32   `json:"points,omitempty"  extensions:"x-omitempty"`
	Reward    int32   `json:"reward,omitempty"  extensions:"x-omitempty"`
	Quantity  int32   `json:"quantity,omitempty"  extensions:"x-omitempty"`
	Price     float32 `json:"price,omitempty"  extensions:"x-omitempty"`
}

type Order struct {
	OrderId     string        `json:"orderId,omitempty"  extensions:"x-omitempty"`
	CustomerId  string        `json:"customerId,omitempty"  extensions:"x-omitempty"`
	Status      Status        `json:"status,omitempty"  extensions:"x-omitempty"`
	AddressId   int32         `json:"addressId,omitempty" default:"nil" extensions:"x-omitempty"`
	TotalPrice  float32       `json:"totalPrice,omitempty" extensions:"x-omitempty"`
	PayedPrice  float32       `json:"payedPrice,omitempty" extensions:"x-omitempty"`
	ProductDesc []ProductDesc `json:"productDesc,omitempty"  extensions:"x-omitempty"`
}

type Status int

const (
	Declined Status = iota + 1
	Pending
	Confirmed
)
