package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	bugsnag "github.com/bugsnag/bugsnag-go"
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

// NewBugsnag return bugsnag notifier
func NewBugsnag() *bugsnag.Notifier {
	errorHandlerConfig := bugsnag.Configuration{
		APIKey:          os.Getenv("BUGSNAG_KEY"),
		ProjectPackages: []string{"main", "github.com/Coderockr/vitrine-social/*"},
	}

	notifier := bugsnag.New(errorHandlerConfig)
	notifier.AutoNotify()

	return notifier
}
