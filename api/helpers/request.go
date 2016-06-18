package helpers

import (
	"net/http"
)

// MustCloseBody is trying to close the req.Body and panics if error occurred
func MustCloseBody(req *http.Request) {
	err := req.Body.Close()
	if err != nil {
		panic(err)
	}
}
