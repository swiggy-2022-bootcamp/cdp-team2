package models
// Customer Response
// @Description Response
type Customer struct {
	Id              string       `json:"customer_id"`
	Firstname       string    `json:"firstname"`
	Lastname        string    `json:"lastname"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Confirm         string    `json:"confirm"`
	Telephone       string    `json:"telephone"`
	Fax             string    `json:"fax"`
	Newsletter		string 		`json:"newsletter"`
	Status          string    `json:"statuss"`
	Approved        string    `json:"approved"`
	Safe            string    `json:"safe"`
	CustomerGroupID int    `json:"customer_group_id"`
	DateAdded       string     `json:"date_added"` 
	Address         []Address `json:"address"`
}