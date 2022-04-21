package dao

type Product struct{
	ProductId	string `json:"product_id"`
	Quality	int `json:"quality"` 
}
type Cart struct{
	Product []Product `json:"product"`
}

type Reward struct{
	CustomerId string `json:"customer_id"`
	Description string `json:"description"`
	Points int   `json:"points"`
}

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
 	UserBalance     string `json:"user_balance"`
	CustomerGroupID string `json:"customer_group_id"`
	DateAdded       string     `json:"date_added"` 
	Agree           string `json:"agree"`
}

type Account2 struct {
	Id              string    `json:"customer_id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Telephone       string `json:"telephone"`
	Cart            Cart    `json:"cart"`
	RewardTotal     Reward    `json:"reward_total"`
	UserBalance     string    `json:"user_balance"`
	CustomerGroupID string `json:"customer_group_id"`
	DateAdded       string     `json:"date_added"` 
}

type Address struct {	 
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Address1  string `json:"address_1"`
	CountryId  string `json:"country_id"`
	Postcode  string `json:"postcode"`
	ZoneId    string `json:"zone_id"`
	Address2  string `json:"address_2"`
	AddressId string `json:"address_id"`
	Default   string `json:"default"`
}


 

//  Authentication -
//  get cartId
//  get reward
//  shipper has to send address
//  transaction  has to check whether
