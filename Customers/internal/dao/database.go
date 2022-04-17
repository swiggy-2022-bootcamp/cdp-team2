package dao

// Customer Account Response
// @Description Account information
type Account struct {
	Id              string    `json:"customer_id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Confirm         string `json:"confirm"`
	Telephone       string `json:"telephone"`
	RewardTotal     string    `json:"reward_total"`
	UserBalance     string    `json:"user_balance"`
	CustomerGroupID string `json:"customer_group_id"`
	Agree           string `json:"agree"`
}

type Account2 struct {
	Id              string    `json:"customer_id"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Telephone       string `json:"telephone"`
	Cart            string    `json:"cart"`
	RewardTotal     string    `json:"reward_total"`
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
	Status          string    `json:"status"`
	Approved        string    `json:"approved"`
	Safe            string    `json:"safe"`
	CustomerGroupID string    `json:"customer_group_id"`
	Cart            string    `json:"cart"`
	RewardTotal     string    `json:"reward_total"`
	UserBalance     string    `json:"user_balance"`
	DateAdded       string     `json:"date_added"` 
	Address         []Address `json:"address"`
}

//  Authentication -
//  get cartId
//  get reward
//  shipper has to send address
//  transaction  has to check whether
