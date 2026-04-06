package routers

import (
	"net/http"

	userHandlers "github.com/Kargozaur/ge/cmd/handlers/user_handlers"
	"github.com/go-chi/chi/v5"
)

func NewUserRouter(h *userHandlers.UserHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/user", func(r chi.Router) {
		r.Post("/register", h.CreateUser())
		r.Post("/login", h.LoginUser())
	})
	return r
}
