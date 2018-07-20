package handlers

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage is a JSON formatter
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleHTTPError formats and returns errors
func HandleHTTPError(w http.ResponseWriter, errno int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errno)
	json.NewEncoder(w).Encode(&ErrorMessage{
		Code:    errno,
		Message: err.Error(),
	})
}
