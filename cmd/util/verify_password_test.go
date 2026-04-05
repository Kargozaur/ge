package util_test

import (
	"testing"

	"github.com/Kargozaur/ge/cmd/util"
)

func TestVerifyPassword(t *testing.T){
	tests := []struct {
		name     string
		password string
		wantErrs int
	}{
		{
			name:     "Valid",
			password: "Abcdef1!",
			wantErrs: 0,
		},
		{
			name:     "Short",
			password: "Ab1!",
			wantErrs: 1,
		},
		{
			name:     "No digit",
			password: "Abcdefg!",
			wantErrs: 1,
		},
		{
			name:     "No upper",
			password: "abcdef1!",
			wantErrs: 1,
		},
		{
			name:     "No special char",
			password: "Abcdef12",
			wantErrs: 1,
		},
		{
			name:     "Multiple errors",
			password: "abc",
			wantErrs: 4,
		},
		{
			name:     "Missing digit and special",
			password: "Abcdefgh",
			wantErrs: 2,
		},
		{
			name:     "Missing upper and digit",
			password: "abcdefg!",
			wantErrs: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			errs := util.VerifyPassword(tt.password)
			if len(errs) != tt.wantErrs {
				t.Errorf("got %d errors, want %d. Errors: %v", len(errs), tt.wantErrs, errs)
			}
		})
	}
}