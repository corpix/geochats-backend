package database

import (
	"github.com/corpix/geochats-backend/config"
	"gopkg.in/mgo.v2"
)

// NewClient creates a new mongo client
func NewClient(conf *config.Config) (*mgo.Database, error) {
	sess, err := mgo.DialWithTimeout(
		conf.DatabaseAddr,
		conf.DatabaseConnectTimeout,
	)
	if err != nil {
		return nil, err
	}

	return sess.DB(conf.DatabaseName), nil
}
