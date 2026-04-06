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
		wr := util.NewJSONWriter(w)
		var userRequest requests.CreateUserRequest
		if err := util.DecodeJson(r, &userRequest); err != nil {
			wr.Write(http.StatusBadRequest, err)
			return
		}
		if userRequest.Email == "" || userRequest.Password == "" {
			wr.WriterError(http.StatusUnprocessableEntity, errors.New("Body must contain email and password fields"))
			return
		}
		user, err := handler.service.CreateUser(&userRequest)
		if err != nil {
			wr.WriterError(http.StatusUnprocessableEntity, err)
			return
		}
		wr.Write(http.StatusOK, user)
	}
}

func (handler *UserHandler) LoginUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wr := util.NewJSONWriter(w)
		var userRequest requests.Login
		if err := util.DecodeJson(r, &userRequest); err != nil {
			wr.WriterError(http.StatusBadRequest, err)
			return
		}
		if userRequest.Email == "" || userRequest.Password == "" {
			wr.WriterError(http.StatusUnprocessableEntity, errors.New("Body must contain email and password fields"))
			return
		}
		token, err := handler.service.VerifyUser(&userRequest)
		if err != nil {
			wr.WriterError(http.StatusNotFound, err)
			return
		}
		cookie := http.Cookie{
			Name:     "access_token",
			Value:    token.AccessToken,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   3600,
			SameSite: http.SameSiteLaxMode,
		}
		wr.SetCookie(&cookie)
		wr.Write(http.StatusOK, token)
	}
}

func (handler *UserHandler) GetUserData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wr := util.NewJSONWriter(w)
		authHeader := r.Header.Get("access_token")
		if authHeader == "" {
			authCookie, err := r.Cookie("access_token")
			if err != nil || authCookie.String() == "" {
				wr.WriterError(http.StatusUnauthorized, errors.New("Unauthorized"))
				return
			}
			authHeader = authCookie.Value
		}
		user, err := handler.service.GetUser(authHeader)
		if err != nil {
			wr.WriterError(http.StatusNotFound, err)
			return
		}
		wr.Write(http.StatusOK, user)
	}
}
