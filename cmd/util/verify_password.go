package util

import (
	"errors"
	"unicode"
)

func VerifyPassword(password string) []error {
	var res []error
	if len(password) < 8 {
		res = append(res, errors.New("Password must be at least 8 characters long"))
	}
	var hasUpper, hasDigit, hasSpecial bool
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case isSpecial(ch):
			hasSpecial = true
		}
	}

	if !hasDigit {
		res = append(res, errors.New("Password must contain at least 1 digit"))
	}
	if !hasUpper {
		res = append(res, errors.New("Passowrd must contain at least 1 upper character"))
	}
	if !hasSpecial {
		res = append(res, errors.New("Password must contain at least 1 special character"))
	}
	return res
}

func isSpecial(ch rune) bool {
	switch ch {
	case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '-', '=':
		return true
	}
	return false
}
