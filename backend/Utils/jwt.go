package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(username string, userId int, isAdmin bool) (string, *ErrorStruct) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":      userId,
			"role":     isAdmin,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", &ErrorStruct{Msg: err.Error(), Code: http.StatusUnauthorized}
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, *ErrorStruct) {
	println("Verifying token ", tokenString)
	//Remove Bearer
	trimmedToken := strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		println("Error :", err.Error())
		return nil, &ErrorStruct{Msg: err.Error(), Code: http.StatusUnauthorized}
	}

	if !token.Valid {
		println("Token invalid")
		return nil, &ErrorStruct{Msg: "Token is invalid", Code: http.StatusUnauthorized}
	}
	println("Token validated continuing...")
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

func GetUserRoleInClaims(ctx *gin.Context) bool {
	tokenClaims, exists := ctx.Get("claims")
	var finalRole bool
	if !exists {
		errorToThrow := &ErrorStruct{Msg: "Invalid token or user", Code: http.StatusInternalServerError}
		ThrowError(ctx, errorToThrow)
		return false
	}
	if claimMap, ok := tokenClaims.(jwt.MapClaims); ok {
		// Par exemple, accéder au champ "sub" dans les claims
		role, ok := claimMap["role"].(bool)
		if !ok {
			ThrowError(ctx, &ErrorStruct{Msg: "User ID not found in claims", Code: http.StatusUnauthorized})
			return false
		}
		finalRole = role
	}
	return finalRole
}
