package geo

import (
	"encoding/json"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/entity"
	storage "github.com/corpix/geochats-backend/storage/geo"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GeoHandlers represents an HTTP handlers to work with geo data
type GeoHandlers struct {
	storage *storage.GeoStorage
}

// GetGeo handles a GET request to the geo endpoint
// And retrieves a geopoints that presented in some area
func (hs *GeoHandlers) GetGeo(resp http.ResponseWriter, req *http.Request) {
	var err error

	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	vars := mux.Vars(req)

	area := map[string]float64{}
	for k, v := range vars {
		area[k], err = strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
	}

	points, err := hs.storage.GetPointsInArea(entity.Area{
		Geo: entity.Geo{
			Latitude:  area["latitude"],
			Longitude: area["longitude"],
		},
		LatitudeDelta:  area["latiudeDelta"],
		LongitudeDelta: area["longitudeDelta"],
	})
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
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	resp.WriteHeader(http.StatusCreated)
}

// Bind mounts API endpoints for geo
func Bind(router *mux.Router) error {
	r := router.PathPrefix("/geo").Subrouter()

	store, err := storage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := GeoHandlers{store}

	r.
		HandleFunc("/{latitude},{longitude},{latitudeDelta},{longitudeDelta}", handlers.GetGeo).
		Methods("GET")

	r.
		HandleFunc("/", handlers.PostGeo).
		Methods("POST")

	return nil
}
