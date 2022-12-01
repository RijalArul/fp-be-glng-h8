package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(pwd string) string {
	salt := 8
	password := []byte(pwd)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
