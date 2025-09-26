package domain

import (
	"time"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Hives []Hive `json:"hives"`
}

type Hive struct {
	ID              string      `json:"id"`
	GeoLocation     GeoLocation `json:"geo_location"`
	Config          string      `json:"config"`
	Description     string      `json:"description"`
	LastServiceDate time.Time   `json:"last_service_date"`
	NextServiceDate time.Time   `json:"next_service_date"`
}

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
