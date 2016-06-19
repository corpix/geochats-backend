package entity

import (
	"gopkg.in/mgo.v2/bson"
)

// Chat represents a chat room
type Chat struct {
	ID      bson.ObjectId `json:"id",bson:"_id,omitempty",hash:"ignore",valid:"ascii,required"`
	PointID bson.ObjectId `json:"pointId",bson:"pointId",valid:",required"`
	Title   string        `json:"title",bson:"title",valid:",required"`
}
