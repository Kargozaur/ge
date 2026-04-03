package services

import (
	"errors"
	"fmt"

	"github.com/Kargozaur/ge/cmd/auth"
	"github.com/Kargozaur/ge/cmd/hasher"
	"github.com/Kargozaur/ge/cmd/models"
	"github.com/Kargozaur/ge/cmd/requests"
	"github.com/Kargozaur/ge/cmd/responses"
	"gorm.io/gorm"
)

type userService struct {
	hasher hasher.PasswordHasher
}

func NewUserService(hasher hasher.PasswordHasher) *userService {
	return &userService{hasher: hasher}
}

func (u *userService) CreateUser(schema *requests.CreateUserRequest, db *gorm.DB) (responses.UserResponse, error) {
	hash, err := u.hasher.Hash(schema.Password)
	if err != nil {
		return responses.UserResponse{}, err
	}
	user := models.ToUserModel(schema.Email, string(hash))
	if err = db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey){
			return responses.UserResponse{}, errors.New("User with this email already exists")
		}
		return responses.UserResponse{}, fmt.Errorf("Failed to create user: %v\n", err)
	}
	userResponse := responses.ToUserResponse(&user)
	return userResponse, nil
}

func (u *userService) VerifyUser(schema *requests.Login, db *gorm.DB, jwtProvider auth.JwtProvider) (responses.Token, error) {
	var userModel models.User
	
	if err := db.Where("email = ?", schema.Email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return responses.Token{}, errors.New("Invalid credentials")
		}
		return responses.Token{}, err
	}
	if !u.hasher.VerifyPwd(schema.Password, userModel.Password) {
		return responses.Token{}, errors.New("Invalid credentials")
	}
	token, err := jwtProvider.CreateAccessToken(userModel.ID)
	if err != nil {
		return responses.Token{}, err
	}
	return responses.NewToken(token), nil
}

func (u *userService) GetUser(jwtToken string, db *gorm.DB, jwtProvider auth.JwtProvider) (responses.UserResponse, error) {
	var userModel models.User
	userId, err := jwtProvider.GetIdFromToken(jwtToken)
	if err != nil {
		return responses.UserResponse{}, err
	}
	if err := db.Where("id = ?", userId).First(&userModel).Error; err != nil {
		return responses.UserResponse{}, err
	}
	userData := responses.ToUserResponse(&userModel)
	return userData, nil
}
