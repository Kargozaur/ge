package util

import (
	"encoding/json"
	"net/http"

	"github.com/Kargozaur/ge/cmd/responses"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
} 

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, responses.ErrorResponse{
		Message: err.Error(),
	})
}