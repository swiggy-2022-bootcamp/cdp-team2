package models

type Reward struct {
	CustomerId  string `json:"customerId" dynamodbav:"customerId"`
	Description string `json:"description" dynamodbav:"description"`
	Points      int32  `json:"points" dynamodbav:"points"`
}
