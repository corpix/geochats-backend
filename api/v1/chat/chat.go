package chat

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/corpix/geochats-backend/api/helpers"
	"github.com/corpix/geochats-backend/config"
	"github.com/corpix/geochats-backend/entity"
	chatStorage "github.com/corpix/geochats-backend/storage/chat"
	messageStorage "github.com/corpix/geochats-backend/storage/message"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const (
	// PathPrefix represents the endpoint prefix to use for API
	PathPrefix = "/chat"
)

// ChatHandlers represents an HTTP handlers that works with chats
type ChatHandlers struct {
	chatStorage    *chatStorage.ChatStorage
	messageStorage *messageStorage.MessageStorage
	router         *mux.Router
}

func (hs *ChatHandlers) validateMessage(message *entity.Message) error {
	_, err := govalidator.ValidateStruct(message)
	return err
}

// GetChat handles a GET request and
// responds with concrete chat by ID from database
func (hs *ChatHandlers) GetChat(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	var err error

	chatID := mux.Vars(req)["chatID"]
	chat, err := hs.chatStorage.GetChat(chatID)
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

// PostChatMessage handles a POST request and posts a message into requested chat
func (hs *ChatHandlers) PostChatMessage(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)
	chatID := mux.Vars(req)["chatID"]

	var err error

	message := &entity.Message{}
	err = json.NewDecoder(req.Body).Decode(message)
	if err != nil {
		panic(err)
	}

	err = hs.validateMessage(message)
	if err != nil {
		helpers.ValidationError(resp, err)
		return
	}

	postedMessage, err := hs.messageStorage.PostMessage(chatID, message)
	if err != nil {
		panic(err)
	}

	retLocation, err := hs.router.
		Get("get-chat-message").
		URL("chatID", chatID, "messageID", postedMessage.ID.Hex())
	if err != nil {
		panic(err)
	}

	resp.WriteHeader(http.StatusCreated)
	resp.Header().Set("location", retLocation.String())

	err = json.NewEncoder(resp).Encode(postedMessage)
	if err != nil {
		panic(err)
	}
}

// GetChatMessage handles a GET request for a single message
func (hs *ChatHandlers) GetChatMessage(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	var err error

	vars := mux.Vars(req)
	chatID := vars["chatID"]
	messageID := vars["messageID"]

	message, err := hs.messageStorage.GetMessage(chatID, messageID)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(resp).Encode(message)
	if err != nil {
		panic(err)
	}
}

// GetChatMessages handles a GET request and
// responds with concrete chat messages
func (hs *ChatHandlers) GetChatMessages(resp http.ResponseWriter, req *http.Request) {
	helpers.JSONResponse(resp)
	defer helpers.MustCloseBody(req)

	var err error

	chatID := mux.Vars(req)["chatID"]
	q := req.URL.Query()
	var (
		limit   int
		startID string
	)
	if limits, ok := q["limit"]; ok && len(limits) > 0 {
		limit, err = strconv.Atoi(limits[0])
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if startIDs, ok := q["startid"]; ok && len(startIDs) > 0 {
		startID = startIDs[0]
	}

	messages, err := hs.messageStorage.GetMessages(chatID, limit, startID)
	err = json.NewEncoder(resp).Encode(messages)
	if err != nil {
		panic(err)
	}
}

// Bind mounts API endpoints for chat
func Bind(router *mux.Router) error {
	chatStore, err := chatStorage.New(config.Get())
	if err != nil {
		return err
	}
	messageStore, err := messageStorage.New(config.Get())
	if err != nil {
		return err
	}

	handlers := ChatHandlers{
		chatStorage:    chatStore,
		messageStorage: messageStore,
		router:         router,
	}

	r := router.PathPrefix(PathPrefix).Subrouter()

	r.
		HandleFunc("/{chatID}", handlers.GetChat).
		Methods("GET").
		Name("get-chat")
	r.
		HandleFunc("/{chatID}/message", handlers.PostChatMessage).
		Methods("POST").
		Name("post-chat-message")
	r.
		HandleFunc("/{chatID}/message/{messageID}", handlers.GetChatMessage).
		Methods("GET").
		Name("get-chat-message")
	r.
		HandleFunc("/{chatID}/messages", handlers.GetChatMessages).
		Methods("GET").
		Name("get-chat-messages")

	return nil
}
