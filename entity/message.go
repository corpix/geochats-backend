package entity

import (
	"github.com/corpix/geochats-backend/pkg/timestamp"
	"gopkg.in/mgo.v2/bson"
)

// Message describes the chat message data
type Message struct {
	User User `json:"user" bson:"user"`
	// FIXME: shit :)
	Name   string              `json:"name" bson:"name" valid:"ascii"`
	ID     bson.ObjectId       `json:"id" bson:"_id" valid:"ascii"`
	ChatID bson.ObjectId       `json:"chatId" bson:"chatId" valid:"ascii"`
	Text   string              `json:"text" bson:"text" valid:",required"`
	Date   timestamp.Timestamp `json:"date" bson:"date" valid:"date"`
}
