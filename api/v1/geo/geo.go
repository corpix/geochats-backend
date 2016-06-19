package geo

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/entity"
	storage "github.com/corpix/geochats-backend/storage/geo"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	// PathPrefix represents the endpoint prefix to use for API
	PathPrefix = "/geo"
)

// GeoHandlers represents an HTTP handlers that works with geo data
type GeoHandlers struct {
	storage *storage.GeoStorage
	router  *mux.Router
}

// GetGeo handles a GET request to the geo endpoint
// And retrieves a geopoints that presented in some area
func (hs *GeoHandlers) GetGeo(resp http.ResponseWriter, req *http.Request) {
	defer helpers.MustCloseBody(req)
	var err error

	helpers.JSONResponse(resp)

	vars := mux.Vars(req)

	areaMap := map[string]float64{}
	for k, v := range vars {
		areaMap[k], err = strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
	}

	area := &entity.Area{
		Geo: entity.Geo{
			Latitude:  areaMap["latitude"],
			Longitude: areaMap["longitude"],
		},
		LatitudeDelta:  areaMap["latitudeDelta"],
		LongitudeDelta: areaMap["longitudeDelta"],
	}

	valid, err := govalidator.ValidateStruct(area)
	if err != nil {
		panic(err)
	}
	if !valid {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	points, err := hs.storage.GetPointsInArea(area)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(resp).Encode(points)
	if err != nil {
		panic(err)
	}
}

// PostGeo handles a POST request to the geo endpoint
// and adds a new geo point to the database
func (hs *GeoHandlers) PostGeo(resp http.ResponseWriter, req *http.Request) {
	defer helpers.MustCloseBody(req)
	var err error

	helpers.JSONResponse(resp)

	userProvidenPoint := &entity.Point{}
	err = json.NewDecoder(req.Body).Decode(userProvidenPoint)
	if err != nil {
		panic(err)
	}

	valid, err := govalidator.ValidateStruct(userProvidenPoint)
	if err != nil {
		panic(err)
	}
	if !valid {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	createdPoint, err := hs.storage.AddPoint(userProvidenPoint)
	if err != nil {
		panic(err)
	}
	if createdPoint == nil {
		resp.WriteHeader(http.StatusConflict)
		return
	}

	resp.WriteHeader(http.StatusCreated)

	retLocation, err := hs.router.Get("get-chat").URL("id", createdPoint.ID.Hex())
	if err != nil {
		panic(err)
	}
	resp.Header().Set("location", retLocation.String())

	err = json.NewEncoder(resp).Encode(createdPoint)
	if err != nil {
		panic(err)
	}
}

// Bind mounts API endpoints for geo
func Bind(router *mux.Router) error {
	store, err := storage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := GeoHandlers{store, router}

	router.
		HandleFunc(PathPrefix, handlers.PostGeo).
		Methods("POST").
		Name("post-geo")

	r := router.PathPrefix(PathPrefix).Subrouter()
	r.
		HandleFunc("/{latitude},{longitude},{latitudeDelta},{longitudeDelta}", handlers.GetGeo).
		Methods("GET").
		Name("get-geo")

	return nil
}
