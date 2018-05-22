package handlers

import (
	"encoding/json"
	"net/http"
)

// HandleHTTPSuccess formats and write body
func HandleHTTPSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleHTTPSuccessNoContent formats and return with no content
func HandleHTTPSuccessNoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
