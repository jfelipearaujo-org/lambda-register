package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsfelipearaujo/lambda-register/src/entities"
)

var (
	signingKey = []byte(os.Getenv("SIGN_KEY"))
)

func CreateJwtToken(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(signingKey)
}