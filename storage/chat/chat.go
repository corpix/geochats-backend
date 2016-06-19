package chat

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/database"
	"github.com/corpix/geochats-backend/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
)

const (
	// CollectionName is a namespace for data storage
	CollectionName = "chat"
)

// ChatStorage represents a struct that works with chat data
type ChatStorage struct {
	database   *mgo.Database
	collection *mgo.Collection
}

// AddChat creates a new chat
func (cs *ChatStorage) AddChat(chat *entity.Chat) (*entity.Chat, error) {
	if chat == nil {
		return nil, errors.New("AddChat: chat is nil")
	}

	var err error

	chatCopy := *chat
	chatCopy.ID = bson.NewObjectId()
	chatCopy.Title = url.QueryEscape(chatCopy.Title)

	err = cs.collection.Insert(chatCopy)
	if err != nil {
		return nil, err
	}

	log.Debugf("AddChat: %+v", chatCopy)
	return &chatCopy, nil
}

// GetChat retrieves a concrete chat from database
func (cs *ChatStorage) GetChat(id string) (*entity.Chat, error) {
	chat := &entity.Chat{}
	log.Debugf("GetChat: %+v", id)
	// FindId is not working, looks like there is a bug
	// Fuck labix!
	err := cs.collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(chat)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return chat, nil
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
