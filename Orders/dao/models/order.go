package models

type ProductDesc struct {
	ProductId int `json:"productId,omitempty"  extensions:"x-omitempty"`
	Points    int `json:"points,omitempty"  extensions:"x-omitempty"`
	Reward    int `json:"reward,omitempty"  extensions:"x-omitempty"`
	Quantity  int `json:"quantity,omitempty"  extensions:"x-omitempty"`
	Price     int `json:"price,omitempty"  extensions:"x-omitempty"`
}

type Order struct {
	OrderId     int           `json:"orderId,omitempty"  extensions:"x-omitempty"`
	CustomerId  int           `json:"customerId,omitempty"  extensions:"x-omitempty"`
	Status      Status        `json:"status,omitempty"  extensions:"x-omitempty"`
	OrderTime   string        `json:"orderTime,omitempty" extensions:"x-omitempty"`
	AddressId   int           `json:"addressId,omitempty" default:"nil" extensions:"x-omitempty"`
	TotalPrice  int           `json:"totalPrice,omitempty" extensions:"x-omitempty"`
	PayedPrice  int           `json:"payedPrice,omitempty" extensions:"x-omitempty"`
	ProductDesc []ProductDesc `json:"productDesc,omitempty"  extensions:"x-omitempty"`
}

type Status string

const (
	Declined Status = Status(iota)
	Pending
	Confirmed
)
