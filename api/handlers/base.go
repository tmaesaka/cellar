package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ErrorInvalidRequest = "invalid_request_error"
)

// ErrorMessage type holds API error related information.
// It is typically serialized to JSON then returned to the client.
type ErrorMessage struct {
	ErrorType string `json:"type"`    // Represents the error classification
	Message   string `json:"message"` // Summary of the error
}

func renderError(w http.ResponseWriter, errorType, message string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")

	errMsg := ErrorMessage{ErrorType: errorType, Message: message}

	if err := json.NewEncoder(w).Encode(errMsg); err != nil {
		log.Print(err)
	}
}

func decodeError(message string) ErrorMessage {
	errMsg := ErrorMessage{}
	json.Unmarshal([]byte(message), &errMsg)
	return errMsg
}
