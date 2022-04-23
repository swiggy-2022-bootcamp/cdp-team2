package models


type Account2 struct {
	Id              string    `json:"customer_id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Telephone       string `json:"telephone"`
	Cart            Cart    `json:"cart"`
	Reward			Reward   `json:"reward"`
	// RewardTotal     Reward    `json:"reward_total"`
	UserBalance     float64    `json:"user_balance"`
	CustomerGroupID int `json:"customer_group_id"`
	DateAdded       string     `json:"date_added"` 
}
 