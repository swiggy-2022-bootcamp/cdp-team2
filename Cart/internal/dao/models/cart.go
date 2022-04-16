package models

type Product struct {
	ProductId string `json:"productId" dynamodbav:"productId"`
	Quantity  int32  `json:"quantity" dynamodbav:"quantity"`
}

type Cart struct {
	CustomerId string    `json:"customerId" dynamodbav:"customerId"`
	Products   []Product `json:"products" dynamodbav:"products"`
}
