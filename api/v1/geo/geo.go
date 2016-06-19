package geo

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/entity"
	chatStorage "github.com/corpix/geochats-backend/storage/chat"
	geoStorage "github.com/corpix/geochats-backend/storage/geo"
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
	geoStorage  *geoStorage.GeoStorage
	chatStorage *chatStorage.ChatStorage
	router      *mux.Router
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

	points, err := hs.geoStorage.GetPointsInArea(area)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(resp).Encode(points)
	if err != nil {
		panic(err)
	}
}

func (hs *GeoHandlers) consumePoint(req *http.Request) *entity.Point {
	var err error

	point := &entity.Point{}

	err = json.NewDecoder(req.Body).Decode(point)
	if err != nil {
		panic(err)
	}

	return point
}

func (hs *GeoHandlers) validatePoint(point *entity.Point) bool {
	valid, err := govalidator.ValidateStruct(point)
	if err != nil {
		panic(err)
	}

	return valid
}

func (hs *GeoHandlers) validateChat(chat *entity.Chat) bool {
	valid, err := govalidator.ValidateStruct(chat)
	if err != nil {
		panic(err)
	}

	return valid
}

func (hs *GeoHandlers) createPoint(point *entity.Point, resp http.ResponseWriter) *entity.Point {
	createdPoint, err := hs.geoStorage.AddPoint(point)
	if err != nil {
		panic(err)
	}

	retLocation, err := hs.router.Get("get-chat").URL("id", createdPoint.ID.Hex())
	if err != nil {
		panic(err)
	}
	resp.Header().Set("location", retLocation.String())

	err = json.NewEncoder(resp).Encode(createdPoint)
	if err != nil {
		panic(err)
	}

	return createdPoint
}

func (hs *GeoHandlers) addChatAtPoint(chat *entity.Chat, point *entity.Point) *entity.Chat {
	chatCopy := *chat
	chatCopy.PointID = point.ID

	createdChat, err := hs.chatStorage.AddChat(&chatCopy)
	if err != nil {
		panic(err)
	}

	return createdChat
}

// PostGeo handles a POST request to the geo endpoint
// and adds a new geo point to the database
func (hs *GeoHandlers) PostGeo(resp http.ResponseWriter, req *http.Request) {
	defer helpers.MustCloseBody(req)
	helpers.JSONResponse(resp)

	userProvidenPoint := hs.consumePoint(req)
	valid := hs.validatePoint(userProvidenPoint)
	if !valid {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	createdPoint := hs.createPoint(userProvidenPoint, resp)

	chat := &entity.Chat{Title: "No name"}
	valid = hs.validateChat(chat)
	if !valid {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	hs.addChatAtPoint(chat, createdPoint)

	resp.WriteHeader(http.StatusCreated)
}

// Bind mounts API endpoints for geo
func Bind(router *mux.Router) error {
	geoStore, err := geoStorage.New(config.Get())
	if err != nil {
		return err
	}
	chatStore, err := chatStorage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := GeoHandlers{
		geoStorage:  geoStore,
		chatStorage: chatStore,
		router:      router,
	}

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
