package auth

import "github.com/google/uuid"

type JwtProvider interface {
	CreateAccessToken(userId uuid.UUID) (string, error)
	VerifyToken(tokenStr string) 		bool
	GetIdFromToken(tokenStr string)		(uuid.UUID, error)	
}