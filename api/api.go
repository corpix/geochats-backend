package api

import (
	"github.com/corpix/geochats-backend/api/v1"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const (
	// Timeout for API requests
	Timeout = 3 * time.Second

	// MIMEType describes a MIME type of requests for API
	MIMEType = "application/json"
)

// New creates new API router
func New() (http.Handler, error) {
	var err error

	router := mux.NewRouter()
	err = v1.Bind(router.PathPrefix("/api").Subrouter())
	if err != nil {
		return nil, err
	}

	timeoutRouter := http.TimeoutHandler(router, Timeout, "Timeout")
	typedRouter := handlers.ContentTypeHandler(timeoutRouter, MIMEType)
	recoveryRouter := handlers.RecoveryHandler()(typedRouter)

	return recoveryRouter, nil
}
