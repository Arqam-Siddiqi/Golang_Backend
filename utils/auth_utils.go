package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwt(_id string) (string, error) {

	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"_id": _id,
		"exp": time.Now().Add(3 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error while signing token: %w", err)
	}

	return signedToken, nil

}
