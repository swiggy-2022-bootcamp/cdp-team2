package models

type Product struct {
	ProductId string `json:"productId" dynamodbav:"productId" example:"12324"`
	Quantity  int32  `json:"quantity" dynamodbav:"quantity" example:"5"`
}

type Cart struct {
	CustomerId string    `json:"customerId" dynamodbav:"customerId" example:"1234"`
	Products   []Product `json:"products" dynamodbav:"products"`
}
