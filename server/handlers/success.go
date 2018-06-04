package handlers

import (
	"encoding/json"
	"net/http"
)

// HandleHTTPSuccess formats and write body
func HandleHTTPSuccess(w http.ResponseWriter, data interface{}, status ...int) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if len(status) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(status[0])
	}
}

// HandleHTTPSuccessNoContent formats and return with no content
func HandleHTTPSuccessNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
