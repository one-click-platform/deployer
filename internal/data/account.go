package data

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
)

type AccountsQ interface {
	New() AccountsQ
	Get() (*Account, error)
	Insert(data Account) (Account, error)
	FilterByEmail(login string) AccountsQ
}

type Account struct {
	ID       int64  `db:"id" structs:"-"`
	Email    string `db:"email" structs:"email"`
	Password []byte `db:"password" structs:"password"`
}

func HashPassword(password []byte) ([]byte, error) {
	// TODO change default cost
	hash, error := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if error != nil {
		return nil, error
	}

	return hash, nil
}

func ComparePassword(password []byte, plainPassword []byte) bool {
	result, error := HashPassword(plainPassword)
	if error != nil {
		return false
	}
	return bytes.Equal(password, result)
}
