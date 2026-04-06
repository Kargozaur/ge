package util

import "net/mail"

func VerifyEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
