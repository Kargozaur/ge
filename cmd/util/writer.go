package util

import (
	"encoding/json"
	"net/http"

	"github.com/Kargozaur/ge/cmd/responses"
)

type JSONWriter struct {
	w http.ResponseWriter
}

func NewJSONWriter(w http.ResponseWriter) *JSONWriter {
	return &JSONWriter{w: w}
}

func (w *JSONWriter) AddHeader(key, value string) *JSONWriter {
	w.w.Header().Add(key, value)
	return w
}

func (w *JSONWriter) SetCookie(cookie *http.Cookie) *JSONWriter {
	http.SetCookie(w.w, cookie)
	return w
}

func (w *JSONWriter) Write(statusCode int, data any) {
	w.w.Header().Set("Content-Type", "application/json")
	w.w.WriteHeader(statusCode)
	json.NewEncoder(w.w).Encode(data)
}

func (w *JSONWriter) WriterError(statusCode int, err error) {
	resp := responses.ErrorResponse{
		Message: err.Error(),
	}
	w.w.Header().Set("Content-Type", "application/json")
	w.w.WriteHeader(statusCode)
	json.NewEncoder(w.w).Encode(resp)
}
