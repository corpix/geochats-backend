package v1

import (
	"github.com/corpix/geochats-backend/api/v1/chat"
	"github.com/corpix/geochats-backend/api/v1/geo"
	"github.com/gorilla/mux"
)

var endpointBinders = []func(*mux.Router) error{
	geo.Bind,
	chat.Bind,
}

// Bind mounts API v1 endpoint
func Bind(router *mux.Router) error {
	var err error
	subrouter := router.PathPrefix("/v1").Subrouter()
	for _, bind := range endpointBinders {
		err = bind(subrouter)
		if err != nil {
			return err
		}
	}
	return nil
}
