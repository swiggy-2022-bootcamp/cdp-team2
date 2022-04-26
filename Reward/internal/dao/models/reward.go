package models

type Reward struct {
	CustomerId  string `json:"customerId" dynamodbav:"customerId" example:"1243"`
	Description string `json:"description" dynamodbav:"description" example:"Reward points in celebration of 50 years of Company ABC"`
	Points      int32  `json:"points" dynamodbav:"points" example:"200"`
}
