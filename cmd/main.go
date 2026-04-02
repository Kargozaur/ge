package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type BaseResponse struct {
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}


func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := BaseResponse{Message: "Default page", When: time.Now().UTC()}
		data, _ := json.Marshal(&msg)
		w.Header().Add("Contenty-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	http.ListenAndServe(":7000", r)
}