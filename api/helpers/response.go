package helpers

import (
	"encoding/json"
	"net/http"
)

type validationError struct {
	Error string `json:"error"`
}

// JSONResponse is a helpers to set correct content type for JSON response
func JSONResponse(resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json; charset=UTF-8")
}

// ValidationError is a helpers to show you whats wrong with validation
func ValidationError(resp http.ResponseWriter, err error) {
	resp.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(resp).Encode(&validationError{err.Error()}); err != nil {
		panic(err)
	}
}
