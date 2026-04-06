package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestTokenManager(t *testing.T) {
	mgr := &TokenKey{SecretKey: []byte("super-secret-key")}
	userID := uuid.New()

	t.Run("CreateAndVerify_Success", func(t *testing.T) {
		tokenStr, err := mgr.CreateAccessToken(userID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		err = mgr.VerifyToken(tokenStr)
		if err != nil {
			t.Fatalf("Expected token to be valid, got error: %v", err)
		}
	})

	t.Run("Verify_InvalidToken", func(t *testing.T) {
		err := mgr.VerifyToken("not.a.valid.token")
		if err == nil {
			t.Error("Error exptected, instead got nil(InvalidToken)")
		}
	})

	t.Run("Verify_WrongKey", func(t *testing.T) {
		tokenStr, _ := mgr.CreateAccessToken(userID)

		wrongMgr := &TokenKey{SecretKey: []byte("another-key")}
		err := wrongMgr.VerifyToken(tokenStr)

		if err == nil {
			t.Error("Error expected, instead got nil(WrongKey)")
		}
	})

	t.Run("Verify_ExpiredToken", func(t *testing.T) {
		expiredClaims := jwt.MapClaims{
			"sub": userID.String(),
			"exp": time.Now().Add(-time.Minute).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
		tokenStr, _ := token.SignedString(mgr.SecretKey)

		err := mgr.VerifyToken(tokenStr)
		if err == nil {
			t.Error("Error expected, instead got nil(ExpiredToken)")
		}
	})
}
