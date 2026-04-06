package userhandlers

import (
	"errors"
	"net/http"

	"github.com/Kargozaur/ge/cmd/auth"
	"github.com/Kargozaur/ge/cmd/hasher"
	"github.com/Kargozaur/ge/cmd/requests"
	"github.com/Kargozaur/ge/cmd/services"
	"github.com/Kargozaur/ge/cmd/util"
	"gorm.io/gorm"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	bcryptHasher := hasher.NewBcryptHasher(10)
	jwt := auth.NewJwtProvider()
	svc := *services.NewUserService(bcryptHasher, jwt, db)
	return &UserHandler{service: svc}
}

func (handler *UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest requests.CreateUserRequest
		if err := util.DecodeJson(r, &userRequest); err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if userRequest.Email == "" || userRequest.Password == "" {
			util.WriteError(w, http.StatusUnprocessableEntity, errors.New("Body must contain email and password fields"))
			return
		}
		user, err := handler.service.CreateUser(&userRequest)
		if err != nil {
			util.WriteError(w, http.StatusUnprocessableEntity, err)
			return
		}
		util.WriteJSON(w, http.StatusOK, user)
	}
}

func (handler *UserHandler) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRequest requests.Login
		if err := util.DecodeJson(r, &userRequest); err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if userRequest.Email == "" || userRequest.Password == "" {
			util.WriteError(w, http.StatusUnprocessableEntity, errors.New("Body must contain email and password fields"))
			return
		}
		token, err := handler.service.VerifyUser(&userRequest)
		if err != nil {
			util.WriteError(w, http.StatusNotFound, err)
		}
		util.WriteJSON(w, http.StatusOK, token)
	}
}
