package utils

import "golang.org/x/crypto/bcrypt"

func HASH(str string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return hash
}
