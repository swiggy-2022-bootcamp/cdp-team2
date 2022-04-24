package api

type SetAddressToOrderRequest struct {
	OrderID   string `json:"orderID,omitempty"`
	AddressId int    `json:"addressId,omitempty"`
}
