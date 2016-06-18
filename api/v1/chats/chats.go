package chats

import (
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

// GetChatHandler handles a GET request and
// responds with concrete chat by ID from database
func GetChatHandler(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)
	resp.Write([]byte(`{}`))
}

// GetChatMessagesHandler handles a GET request and
// responds with concrete chat messages
func GetChatMessagesHandler(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)
	resp.Write([]byte(`[]`))
}

// Bind mounts API endpoints for chats
func Bind(router *mux.Router) error {
	r := router.PathPrefix("/chats").Subrouter()

	r.
		HandleFunc("/{id}", GetChatHandler).
		Methods("GET")

	r.
		HandleFunc("/{id}/messages", GetChatMessagesHandler).
		Methods("GET")

	return nil
}
