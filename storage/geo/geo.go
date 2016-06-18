package geo

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
	CollectionName = "geo"

	// MaxPointsInArea limits the maximum points in area to get from database
	MaxPointsInArea = 100
)

// GeoStorage represents a struct that works with geo data storage
type GeoStorage struct {
	database   *mgo.Database
	collection *mgo.Collection
}

// AddPoint stores a new Point in the storage
func (gs *GeoStorage) AddPoint(point *entity.Point) (*entity.Point, error) {
	if point == nil {
		return nil, errors.New("AddPoint: point is nil")
	}

	point.ID = bson.NewObjectId()
	point.Title = url.QueryEscape(point.Title)

	err := gs.collection.Insert(point)
	if err != nil {
		return nil, err
	}

	log.Debugf("AddPoint: %+v", point)
	return point, nil
}

// GetPointsInArea retrieves point in the specified area from storage
func (gs *GeoStorage) GetPointsInArea(area *entity.Area) ([]entity.Point, error) {
	if area == nil {
		return nil, errors.New("GetPointsInArea: area is nil")
	}
	log.Debugf(
		"geo.latitude gt %+v lt %+v geo.longitude gt %+v lt %+v",
		area.Latitude-area.LatitudeDelta,
		area.Latitude+area.LatitudeDelta,
		area.Longitude-area.LongitudeDelta,
		area.Longitude+area.LongitudeDelta,
	)
	iter := gs.collection.Find(
		bson.M{
			"geo.latitude": bson.M{
				"$gt": area.Latitude - area.LatitudeDelta,
				"$lt": area.Latitude + area.LatitudeDelta,
			},
			"geo.longitude": bson.M{
				"$gt": area.Longitude - area.LongitudeDelta,
				"$lt": area.Longitude + area.LongitudeDelta,
			},
		},
	).Limit(MaxPointsInArea).Iter()

	point := entity.Point{}
	points := []entity.Point{}
	for iter.Next(&point) {
		points = append(points, point)
	}

	return points, nil
}

// New creates a new geo storage instance
func New(conf *config.Config) (*GeoStorage, error) {
	db, err := database.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &GeoStorage{
		database:   db,
		collection: db.C(CollectionName),
	}, nil
}
