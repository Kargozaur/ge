package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenKey struct {
	SecretKey []byte
}

func NewSecretKey() *TokenKey {
	return &TokenKey{SecretKey: []byte(os.Getenv("SECRET_KEY"))}
}


func (t *TokenKey) CreateAccessToken(userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId.String(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString(t.SecretKey)
}

func (t *TokenKey) VerifyToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error){
		return t.SecretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("Token expired")
	}
	return nil
}