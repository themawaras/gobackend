package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecretKeys = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	Username string `json: "username"`
	jwt.StandardClaims
}

func NewToken(username string) *claims {
	return &claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return tokens.SignedString(mySecretKeys)
}

func CheckToken(token string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretKeys), nil
	})

	if err != nil {
		return false, err
	}

	return tokens.Valid, nil
}
