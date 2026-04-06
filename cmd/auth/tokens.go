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

func NewJwtProvider() *TokenKey {
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
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
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

func (t *TokenKey) GetIdFromToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Failed to decode access token")
		}
		return t.SecretKey, nil
	})
	if err != nil {
		return uuid.Nil, nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			return uuid.Nil, errors.New("Failed to find sub parameter in the jwt token")
		}
		userId, err := uuid.Parse(sub)
		if err != nil {
			return uuid.Nil, errors.New("Failed to parse UUID")
		}
		return userId, nil
	}
	return uuid.Nil, errors.New("Invalid token")
}
