package models
type Reward struct{
	CustomerId string `json:"customer_id"`
	Description string `json:"description"`
	Points int   `json:"points"`
}
