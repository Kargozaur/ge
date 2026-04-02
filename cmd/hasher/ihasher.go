package hasher

type PasswordHasher interface {
	Hash(password string) (string, error)
	VerifyPwd(password, hashedPwd string) bool
}

