package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	println("Verifying token ", tokenString)
	//Remove Bearer
	trimmedToken := strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		println("Error :", err.Error())
		return nil, err
	}

	if !token.Valid {
		println("Token invalid")
		return nil, errors.New("Unauthorized")
	}
	println("Token validated continuing...")
	return token, nil
}
