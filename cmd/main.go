package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Kargozaur/ge/cmd/config"
	userhandlers "github.com/Kargozaur/ge/cmd/handlers/user_handlers"
	"github.com/Kargozaur/ge/cmd/routers"
	"github.com/go-chi/chi/v5"
)

type BaseResponse struct {
	Message string    `json:"message"`
	When    time.Time `json:"when"`
}

func main() {
	db := config.DbConf()
	userHandler := userhandlers.NewUserHandler(db)
	userRouter := routers.NewUserRouter(userHandler)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := BaseResponse{Message: "Default page", When: time.Now().UTC()}
		data, _ := json.Marshal(&msg)
		w.Header().Add("Contenty-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	r.Mount("/", userRouter)
	http.ListenAndServe(":7000", r)
}
