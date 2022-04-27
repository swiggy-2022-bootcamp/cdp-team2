package models
type Address struct {	 
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	AddressLine1  string `json:"address_1"`
	AddressLine2  string `json:"address_2"`
	City      string `json:"city"`
	PostCode  int32 `json:"postcode"`
	CountryCode  string `json:"country_id"`
}