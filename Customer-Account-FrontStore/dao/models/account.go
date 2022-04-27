package models

// Customer Account Response
// @Description Account information
type Account struct {
	Id              string `json:"customer_id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Confirm         string `json:"confirm"`
	Telephone       string `json:"telephone"`
 	UserBalance     float64 `json:"user_balance"`
	CustomerGroupID int `json:"customer_group_id"`
	DateAdded       string     `json:"date_added"` 
	Agree           string `json:"agree"`
}