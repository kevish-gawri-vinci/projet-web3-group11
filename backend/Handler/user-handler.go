package handler

import (
	request "backend/Request"
	service "backend/Service"
	utils "backend/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserHandler(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRequest request.UserRequest
		bindError := c.ShouldBind(&userRequest)
		if bindError != nil {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: "Bad input", Code: http.StatusBadRequest})
			return
		}
		isPasswordOK, err := utils.CheckPassword(userRequest.Password)
		if !isPasswordOK {
			utils.ThrowError(c, err)
			return
		}
		err = userService.AddUser(userRequest)
		if err != nil {
			utils.ThrowError(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{})
	}
}

func LoginHandler(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputRequest request.UserRequest

		if err := c.BindJSON(&inputRequest); err != nil {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: "Bad input", Code: http.StatusBadRequest})
			return
		}
		token, errorToThrow := userService.Login(inputRequest)

		if errorToThrow != nil {
			utils.ThrowError(c, errorToThrow)
			return
		}
		c.Header("Authorization", "Bearer "+token)
		c.JSON(http.StatusAccepted, gin.H{})
	}
}

func GetRoleHandler(userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//User is already authentified => authmiddleware
		userId := utils.GetUserIdInClaims(ctx)
		if userId == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "Error: no ID found in claims"})
		}
		response, err := userService.GetUserRole(userId)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(200, gin.H{
			"response": response,
		})
	}
}
