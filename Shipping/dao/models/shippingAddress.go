package models

type ShippingAddress struct {
	AddressId    int    `json:"addressId,omitempty"  extensions:"x-omitempty"`
	CustomerId   int    `json:"customerId,omitempty"  extensions:"x-omitempty"`
	FirstName    string `json:"firstName,omitempty" extensions:"x-omitempty"`
	LastName     string `json:"lastName,omitempty" extensions:"x-omitempty"`
	City         string `json:"city,omitempty" extensions:"x-omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty" extensions:"x-omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty" extensions:"x-omitempty"`
	CountryCode  string `json:"countryCode,omitempty" extensions:"x-omitempty"`
	PostCode     int    `json:"postCode,omitempty" extensions:"x-omitempty"`
}
