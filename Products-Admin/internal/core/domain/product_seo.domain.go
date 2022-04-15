package domain

type ProductSeoUrl struct {
	Keyword    string `json:"keyword" bson:"keyword"`
	LanguageID int    `json:"language_id" bson:"langauge_id"`
	StoreID    int    `json:"store_id" bson:"store_id"`
}
