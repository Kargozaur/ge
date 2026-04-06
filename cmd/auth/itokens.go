package auth

import "github.com/google/uuid"

type JwtProvider interface {
	CreateAccessToken(userId uuid.UUID) (string, error)
	VerifyToken(tokenStr string) error
	GetIdFromToken(tokenStr string) (uuid.UUID, error)
}
