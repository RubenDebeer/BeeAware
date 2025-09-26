package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserID = primitive.ObjectID

type User struct {
	ID        UserID `bson:"_id,omitempty" json:"id"`
	Name      string `bson:"name" json:"name"`
	HiveCount int    `bson:"hivecount" json:"hivecount"`
}
