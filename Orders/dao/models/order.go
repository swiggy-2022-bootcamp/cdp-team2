package models

type ProductDesc struct {
	ProductId int32 `json:"productId,omitempty"  extensions:"x-omitempty"`
	Points    int32 `json:"points,omitempty"  extensions:"x-omitempty"`
	Reward    int32 `json:"reward,omitempty"  extensions:"x-omitempty"`
	Quantity  int32 `json:"quantity,omitempty"  extensions:"x-omitempty"`
	Price     int32 `json:"price,omitempty"  extensions:"x-omitempty"`
}

type Order struct {
	OrderId     int32         `json:"orderId,omitempty"  extensions:"x-omitempty"`
	CustomerId  string        `json:"customerId,omitempty"  extensions:"x-omitempty"`
	Status      Status        `json:"status,omitempty"  extensions:"x-omitempty"`
	AddressId   int32         `json:"addressId,omitempty" default:"nil" extensions:"x-omitempty"`
	TotalPrice  int32         `json:"totalPrice,omitempty" extensions:"x-omitempty"`
	PayedPrice  int32         `json:"payedPrice,omitempty" extensions:"x-omitempty"`
	ProductDesc []ProductDesc `json:"productDesc,omitempty"  extensions:"x-omitempty"`
}

type Status int32

const (
	Declined Status = iota + 1
	Pending
	Confirmed
)
