package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func PasswordsEqual(password []byte, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword(password, []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
