package message

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/database"
	"github.com/corpix/geochats-backend/entity"
	"github.com/corpix/geochats-backend/pkg/timestamp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"time"
)

const (
	// CollectionName is a namespace for data storage
	CollectionName = "message"

	// MaxMessagesPerRetrieval limits the maximum amount of messages that
	// could be retrieved via single query to the database
	MaxMessagesPerRetrieval = 50
)

// MessageStorage represents a struct that works with message data
type MessageStorage struct {
	database   *mgo.Database
	collection *mgo.Collection
}

// PostMessage posts a new message into database tagged with chat id
func (ms *MessageStorage) PostMessage(chatID string, message *entity.Message) (*entity.Message, error) {
	if message == nil {
		return nil, errors.New("PostMessage: message is nil")
	}

	var err error

	messageCopy := *message
	messageCopy.ID = bson.NewObjectId()
	messageCopy.Date = timestamp.Timestamp(time.Now())
	messageCopy.ChatID = bson.ObjectIdHex(chatID)
	// FIXME: Find good way to escape this to work with RN
	//messageCopy.Text = url.QueryEscape(message.Text)

	err = ms.collection.Insert(messageCopy)
	if err != nil {
		return nil, err
	}

	log.Debugf("PostMessage: %+v <- %+v", chatID, message)
	return &messageCopy, nil
}

// GetMessage retrieves a message from database
func (ms *MessageStorage) GetMessage(chatID string, messageID string) (*entity.Message, error) {
	var err error

	message := &entity.Message{}
	err = ms.collection.Find(
		bson.M{
			"_id":    bson.ObjectIdHex(messageID),
			"chatId": bson.ObjectIdHex(chatID),
		},
	).One(message)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return message, nil
}

// GetMessages retrieves chat messages with specified limit and offset from database
func (ms *MessageStorage) GetMessages(chatID string, limit int, startID string) ([]entity.Message, error) {
	if limit == 0 || limit > MaxMessagesPerRetrieval {
		limit = MaxMessagesPerRetrieval
	}

	messages := []entity.Message{}
	iter := ms.collection.
		Find(bson.M{"chatId": bson.ObjectIdHex(chatID)}).
		Sort("date").
		Limit(limit).Iter()
	message := &entity.Message{}
	for iter.Next(message) {
		messages = append(messages, *message)
	}

	return messages, nil
}

// New creates a new geo storage instance
func New(conf *config.Config) (*MessageStorage, error) {
	db, err := database.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &MessageStorage{
		database:   db,
		collection: db.C(CollectionName),
	}, nil
}
