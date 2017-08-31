package routes

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage is a JSON formatter
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleHttpError formats and returns errors
func HandleHttpError(w http.ResponseWriter, errno int, err error) {
	w.WriteHeader(errno)
	json.NewEncoder(w).Encode(&ErrorMessage{
		Code:    errno,
		Message: err.Error(),
	})
}
