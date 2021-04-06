package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		return nil, error
	}

	return hash, nil
}

func PasswordsEqual(password []byte, plainPassword string) bool {
	error := bcrypt.CompareHashAndPassword(password, []byte(plainPassword))
	if error != nil {
		return false
	}
	return true
}
