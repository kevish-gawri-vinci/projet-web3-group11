package utils

import (
	"net/http"
	"regexp"
	"unicode"
)

/*
* Function checking the password in parameter
 */

func CheckPassword(password string) (bool, *ErrorStruct) {
	println(len(password) > 7)
	if len(password) < 8 {
		return false, &ErrorStruct{Msg: "Password must contain at least 8 characters", Code: http.StatusForbidden}
	}
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false, &ErrorStruct{Msg: "Password must contain a number", Code: http.StatusForbidden}
	}
	if !containsUpper(password) {
		return false, &ErrorStruct{Msg: "Password must contain an uppercase character", Code: http.StatusForbidden}
	}
	return true, nil
}

func containsUpper(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}
