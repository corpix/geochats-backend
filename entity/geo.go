package entity

import (
	"gopkg.in/mgo.v2/bson"
)

// Geo represents a geo data
type Geo struct {
	Latitude  float64 `json:"latitude",bson:"latitude",valid:"latitude,required"`
	Longitude float64 `json:"longitude",bson:"longitude",valid:"longitude,required"`
}

// Point represents a uniquely addressable point on the map
// Eg, the point that has an ID as a primary key and some data associated with it(title, ...)
type Point struct {
	ID bson.ObjectId `json:"id",bson:"_id,omitempty",hash:"ignore",valid:"ascii,required"`
	Geo
}

// Area represents an area on the map
type Area struct {
	Geo
	LatitudeDelta  float64 `json:"latitudeDelta",bson:"latitudeDelta",valid:"numeric,required"`
	LongitudeDelta float64 `json:"longitudeDelta",bson:"longitudeDelta",valid:"numeric,required"`
}
