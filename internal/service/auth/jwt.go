package auth

import (
	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/one-click-platform/deployer/internal/data"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"strconv"
	"time"
)

type JWToken interface {
	ReadConfig() JWTokenCfg
}

type jwtoken struct {
	getter kv.Getter
	once   comfig.Once
}

func NewJWToken(getter kv.Getter) JWToken {
	return &jwtoken{
		getter: getter,
	}
}

type JWTokenCfg struct {
	Secret string `fig:"secret,required"`
}

type PayloadJWT struct {
	exp float64 `fig:"exp"`
	sub string  `fig:"sub"`
}

func (d *jwtoken) ReadConfig() JWTokenCfg {
	var cfg JWTokenCfg
	err := figure.Out(&cfg).
		From(kv.MustGetStringMap(d.getter, "jwt")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return cfg
}

// TODO rework structure
func (cfg JWTokenCfg) CreateToken(account *data.Account) (string, error) {
	signingKey := []byte(cfg.Secret)

	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Minute * 15).Unix())),
		Subject:   strconv.FormatInt(account.ID, 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func (cfg JWTokenCfg) VerifyToken(t string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}
