package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductSeoUrl struct {
	Keyword    string             `json:"keyword" bson:"keyword"`
	LanguageID primitive.ObjectID `json:"language_id" bson:"langauge_id"`
	StoreID    primitive.ObjectID `json:"store_id" bson:"store_id"`
}
