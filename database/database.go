package database

import (
	"github.com/corpix/geochats-backend/config"
	"gopkg.in/redis.v3"
)

// NewClient creates a new redis client
func NewClient(conf *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: conf.DatabaseAddr,
		// FIXME: Whats about security man?
		Password: "",
		// FIXME: Customizable database?
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
