package helpers

import (
	"net/http"
)

// JSONResponse is a helpers to set correct content type for JSON response
func JSONResponse(resp http.ResponseWriter) {
	resp.Header().Set("content-type", "application/json; charset=UTF-8")
}
