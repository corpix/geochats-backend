package chat

import (
	"encoding/json"
	//"github.com/asaskevich/govalidator"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	chatStorage "github.com/corpix/geochats-backend/storage/chat"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	// PathPrefix represents the endpoint prefix to use for API
	PathPrefix = "/chat"
)

// ChatHandlers represents an HTTP handlers that works with chats
type ChatHandlers struct {
	chatStorage *chatStorage.ChatStorage
}

// GetChat handles a GET request and
// responds with concrete chat by ID from database
func (hs *ChatHandlers) GetChat(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	var err error

	id := mux.Vars(req)["id"]
	chat, err := hs.chatStorage.GetChat(id)
	if err != nil {
		panic(err)
	}
	if chat == nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(resp).Encode(chat)
	if err != nil {
		panic(err)
	}
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
	chatStore, err := chatStorage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := ChatHandlers{
		chatStorage: chatStore,
	}

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
