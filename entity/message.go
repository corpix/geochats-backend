package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Message describes the chat message data
type Message struct {
	User   User          `json:"user" bson:"user"`
	ID     bson.ObjectId `json:"id" bson:"_id" valid:"ascii"`
	ChatID bson.ObjectId `json:"chatId" bson:"chatId" valid:"ascii"`
	Text   string        `json:"text" bson:"text" valid:",required"`
	Date   time.Time     `json:"date" bson:"date" valid:"date"`
}
