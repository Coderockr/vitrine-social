package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

// HandleHTTPSuccessImage formats and write image
func HandleHTTPSuccessImage(w http.ResponseWriter, data io.ReadCloser) {
	buffer := new(bytes.Buffer)
	_, _ = io.Copy(buffer, data)

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(buffer.Bytes()); err != nil {
		HandleHTTPError(w, http.StatusInternalServerError, err)
		return
	}
}
