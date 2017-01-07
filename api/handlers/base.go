package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
)

const (
	ErrorApi            = "api_error"
	ErrorInvalidRequest = "invalid_request_error"
)

// ErrorMessage type holds API error related information.
// It is typically serialized to JSON then returned to the client.
type ErrorMessage struct {
	ErrorType string `json:"type"`    // Represents the error classification
	Message   string `json:"message"` // Summary of the error
}

// BadRequest renders a 400 bad request response.
func BadRequest(w http.ResponseWriter, errorType, message string) {
	errMsg := ErrorMessage{ErrorType: errorType, Message: message}
	renderError(w, http.StatusBadRequest, &errMsg)
}

// NotFound renders a 404 not found response.
func NotFound(w http.ResponseWriter) {
	renderError(w, http.StatusNotFound, nil)
}

func renderError(w http.ResponseWriter, status int, errMsg *ErrorMessage) {
	w.WriteHeader(status)

	if errMsg != nil {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(errMsg); err != nil {
			log.Print(err)
		}
	}
}

func decodeError(message string) ErrorMessage {
	errMsg := ErrorMessage{}
	json.Unmarshal([]byte(message), &errMsg)
	return errMsg
}

func repoPath(datadir, repoName string) string {
	return path.Join(datadir, "repos/", repoName)
}
