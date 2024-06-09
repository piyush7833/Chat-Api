package helpers

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standard error response.
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// SuccessResponse represents a standard success response.
type SuccessResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// JSON sends a JSON response to the client.
func JSON(w http.ResponseWriter, status int, payload interface{}, message string) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// Error sends a JSON error response to the client.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorResponse{
		Status:  status,
		Message: message,
	}, message)
}

// Success sends a JSON success response to the client.
func Success(w http.ResponseWriter, status int, data interface{}, message string) {
	JSON(w, status, SuccessResponse{
		Status:  status,
		Data:    data,
		Message: message,
	}, message)
}
