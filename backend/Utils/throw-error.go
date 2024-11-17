package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ThrowError(ctx *gin.Context, err *ErrorStruct) {
	ctx.Set("statusCode", err.Code)
	ctx.Error(errors.New(err.Msg))
}

type ErrorStruct struct {
	Msg  string
	Code int
}

func (e *ErrorStruct) Error() string {
	return e.Msg
}
