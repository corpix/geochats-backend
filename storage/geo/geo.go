package geo

import (
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/database"
	"github.com/corpix/geochats-backend/entity"
	"gopkg.in/redis.v3"
)

// GeoStorage represents a struct that works with geo data storage
type GeoStorage struct {
	database *redis.Client
}

// GetPointsInArea retrieves point in the specified area from storage
func (gs *GeoStorage) GetPointsInArea(area entity.Area) ([]entity.Point, error) {
	return []entity.Point{
		{
			ID:    "deadbeef1",
			Title: "Поговори со мной",
			Geo: entity.Geo{
				Latitude:  55.752141,
				Longitude: 37.6143558,
			},
		},
		{
			ID:    "deadbeef2",
			Title: "Ололо",
			Geo: entity.Geo{
				Latitude:  55.752141,
				Longitude: 37.6143558,
			},
		},
		{
			ID:    "deadbeef3",
			Title: "Я в домике",
			Geo: entity.Geo{
				Latitude:  55.75111,
				Longitude: 37.613486,
			},
		},
		{
			ID:    "deadbeef4",
			Title: "Пыщ пыщ",
			Geo: entity.Geo{
				Latitude:  55.751038,
				Longitude: 37.609066,
			},
		},
		{
			ID:    "deadbeef5",
			Title: "Пепяка ололо, Онотоле попячься!",
			Geo: entity.Geo{
				Latitude:  55.755216,
				Longitude: 37.611684,
			},
		},
		{
			ID:    "deadbeef6",
			Title: "test",
			Geo: entity.Geo{
				Latitude:  55.755192,
				Longitude: 37.619923,
			},
		},
	}, nil
}

// New creates a new geo storage instance
func New(conf *config.Config) (*GeoStorage, error) {
	db, err := database.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &GeoStorage{db}, nil
}
