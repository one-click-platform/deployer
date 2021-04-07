package auth

import (
	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/one-click-platform/deployer/internal/data"
	"strconv"
	"time"
)

func CreateToken(account data.Account) (string, error) {
	signingKey := []byte("aludfhalisfhasd312f12kbk34")

	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Minute * 15).Unix())),
		Subject:   strconv.FormatInt(account.ID, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func VerifyToken(t string) error {
	_, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte("aludfhalisfhasd312f12kbk34"), nil
	})
	if err != nil {
		return err
	}

	return nil
}
