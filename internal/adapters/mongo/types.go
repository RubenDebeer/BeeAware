package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//
// Mongo persistence types (DTOs)
//

type MongoUser struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Hives []MongoHive        `bson:"hives"`
}

type MongoHive struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	GeoLocation     MongoGeoLocation   `bson:"geo_location"`
	Config          string             `bson:"config"`
	Description     string             `bson:"description"`
	LastServiceDate time.Time          `bson:"last_service_date"`
	NextServiceDate time.Time          `bson:"next_service_date"`
}

type MongoGeoLocation struct {
	Lat float64 `bson:"lat"`
	Lon float64 `bson:"lon"`
}
