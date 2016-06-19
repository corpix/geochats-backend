package chat

import (
	//"encoding/json"
	//"github.com/asaskevich/govalidator"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	storage "github.com/corpix/geochats-backend/storage/chat"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	// PathPrefix represents the endpoint prefix to use for API
	PathPrefix = "/chat"
)

// ChatHandlers represents an HTTP handlers that works with chats
type ChatHandlers struct {
	storage *storage.ChatStorage
}

// GetChat handles a GET request and
// responds with concrete chat by ID from database
func (hs *ChatHandlers) GetChat(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)
	//vars := mux.Vars(req)
	//id := vars["id"]
	return
}

// GetChatMessages handles a GET request and
// responds with concrete chat messages
func (hs *ChatHandlers) GetChatMessages(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)
	resp.Write([]byte(`[]`))
}

// Bind mounts API endpoints for chat
func Bind(router *mux.Router) error {
	store, err := storage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := ChatHandlers{store}

	r := router.PathPrefix(PathPrefix).Subrouter()

	r.
		HandleFunc("/{id}", handlers.GetChat).
		Methods("GET").
		Name("get-chat")

	r.
		HandleFunc("/{id}/messages", handlers.GetChatMessages).
		Methods("GET").
		Name("get-chat-messages")

	return nil
}
