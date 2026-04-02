package hasher_test

import (
	"testing"

	"github.com/Kargozaur/ge/cmd/hasher"
)

func TestBcryptHasher(t *testing.T) {
	h := hasher.NewBcryptHasher(10)
	tests := []struct {
		name     string
		password string
	}{
		{
			name:     "Casual",
			password: "my-secret-password",
		},
		{
			name:     "Short pass",
			password: "123",
		},
		{
			name:     "Empty pass",
			password: "",
		},
		{
			name:     "Password with special symbols",
			password: "password!@#$%^&*()",
		},
	}
	for _, tt := range tests {
		hash, err := h.Hash(tt.password)
		if err != nil {
			t.Errorf("Failed to hash %v\n", err)
		}
		if hash == "" {
			t.Error("Hash can't be an empty string")
		}
		if !h.VerifyPwd(tt.password, hash) {
			t.Errorf("Failed to verify password %v\n", tt.password)
		}

	}
}

func TestSalt(t *testing.T) {
	h := hasher.NewBcryptHasher(10)
	pwd := "myCoolPassword"
	hash1, _ := h.Hash(pwd)
	hash2, _ := h.Hash(pwd)
	if hash1 == hash2 {
		t.Errorf("Hasher generated the same hash")
	}
}