package chat

import (
	//"errors"
	//log "github.com/Sirupsen/logrus"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/database"
	"github.com/corpix/geochats-backend/entity"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"net/url"
)

const (
	// CollectionName is a namespace for data storage
	CollectionName = "chat"

	// MaxMessagesPerRetrieval limits the maximum amount of messages that
	// could be retrieved via single query to the database
	MaxMessagesPerRetrieval = 50
)

// ChatStorage represents a struct that works with chat data
type ChatStorage struct {
	database   *mgo.Database
	collection *mgo.Collection
}

// GetChatMessages
func GetChatMessages(id string, limit int) ([]entity.Message, error) {
	return nil, nil
}

// New creates a new geo storage instance
func New(conf *config.Config) (*ChatStorage, error) {
	db, err := database.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &ChatStorage{
		database:   db,
		collection: db.C(CollectionName),
	}, nil
}
