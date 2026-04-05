package routers

import (
	"net/http"

	userHandlers "github.com/Kargozaur/ge/cmd/handlers/user_handlers"
	"github.com/go-chi/chi/v5"
)

type RouterDeps struct {
	UserHandler *userHandlers.UserHandler
}

func NewUserRouter(h RouterDeps) http.Handler{
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router){
		r.Post("/register", h.UserHandler.CreateUser())
		r.Post("/login", h.UserHandler.LoginUser())
	})
	return r
}