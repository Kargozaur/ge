package services

import (
	"errors"
	"fmt"

	"github.com/Kargozaur/ge/cmd/auth"
	"github.com/Kargozaur/ge/cmd/hasher"
	"github.com/Kargozaur/ge/cmd/models"
	"github.com/Kargozaur/ge/cmd/requests"
	"github.com/Kargozaur/ge/cmd/responses"
	"github.com/Kargozaur/ge/cmd/util"
	"gorm.io/gorm"
)

type UserService struct {
	hasher      hasher.PasswordHasher
	jwtProvider auth.JwtProvider
	db          *gorm.DB
}

func NewUserService(hasher hasher.PasswordHasher, jwtProvider auth.JwtProvider, db *gorm.DB) *UserService {
	return &UserService{hasher: hasher, jwtProvider: jwtProvider, db: db}
}

func (u *UserService) CreateUser(schema *requests.CreateUserRequest) (responses.UserResponse, error) {
	if !util.VerifyEmail(schema.Email) {
		return responses.UserResponse{}, errors.New("Invalid email")
	}
	if errs := util.VerifyPassword(schema.Password); len(errs) > 0 {
		return responses.UserResponse{}, errors.Join(errs...)
	}
	hash, err := u.hasher.Hash(schema.Password)
	if err != nil {
		return responses.UserResponse{}, err
	}
	user := models.ToUserModel(schema.Email, string(hash))
	if err = u.db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return responses.UserResponse{}, errors.New("User with this email already exists")
		}
		return responses.UserResponse{}, fmt.Errorf("Failed to create user: %v\n", err)
	}
	userResponse := responses.ToUserResponse(&user)
	return userResponse, nil
}

func (u *UserService) VerifyUser(schema *requests.Login) (responses.Token, error) {
	if !util.VerifyEmail(schema.Email) {
		return responses.Token{}, fmt.Errorf("Email %s is not a valid email", schema.Email)
	}
	if errs := util.VerifyPassword(schema.Password); len(errs) > 0 {
		return responses.Token{}, fmt.Errorf("Bad password format. %s", errors.Join(errs...))
	}
	var userModel models.User
	if err := u.db.Where("email = ?", schema.Email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.Token{}, errors.New("Invalid credentials")
		}
		return responses.Token{}, err
	}
	if !u.hasher.VerifyPwd(schema.Password, userModel.Password) {
		return responses.Token{}, errors.New("Invalid credentials")
	}
	token, err := u.jwtProvider.CreateAccessToken(userModel.ID)
	if err != nil {
		return responses.Token{}, err
	}
	return responses.NewToken(token), nil
}

func (u *UserService) GetUser(jwtToken string) (responses.UserResponse, error) {
	var userModel models.User
	userId, err := u.jwtProvider.GetIdFromToken(jwtToken)
	if err != nil {
		return responses.UserResponse{}, err
	}
	if err := u.db.Where("id = ?", userId).First(&userModel).Error; err != nil {
		return responses.UserResponse{}, err
	}
	userData := responses.ToUserResponse(&userModel)
	return userData, nil
}
