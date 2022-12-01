package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(pwd string) string {
	salt := 8
	password := []byte(pwd)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}
