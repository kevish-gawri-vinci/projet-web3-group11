package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      userId,
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
	// println("Retrieving userId from token ")
	// var userId string
	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	sub, ok := claims["sub"].(string)
	// 	if !ok {
	// 		err := errors.New("Invalid token payload")
	// 		return nil, err, ""
	// 	}

	// 	// Put the sub in the variable to return
	// 	userId = sub
	// } else {
	// 	err := errors.New("Invalid token")
	// 	return nil, err, ""
	// }
	return token, nil
}

func GetUserIdInClaims(ctx *gin.Context) int {
	tokenClaims, exists := ctx.Get("claims")
	var finalId int
	if !exists {
		errorToThrow := &ErrorStruct{Msg: "Invalid token or user", Code: http.StatusInternalServerError}
		ThrowError(ctx, errorToThrow)
		return 0
	}
	if claimMap, ok := tokenClaims.(jwt.MapClaims); ok {
		// Par exemple, accéder au champ "sub" dans les claims
		userID, ok := claimMap["sub"].(float64)
		if !ok {
			ThrowError(ctx, &ErrorStruct{Msg: "User ID not found in claims", Code: http.StatusUnauthorized})
			return 0
		}
		userIDInt := int(userID)
		println(userIDInt)
		finalId = userIDInt
	}
	return finalId
}
