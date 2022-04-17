package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id",omitempty`
	Role   string             `bson:"role"`
	Tokens []string           `bson:"tokens"`
}
