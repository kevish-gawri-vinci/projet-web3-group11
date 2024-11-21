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
		if err := c.ShouldBind(&userRequest); err != nil {
			c.Error(err)
			return
		}
		user, err := userService.AddUser(userRequest)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"password": user.Password,
		})
	}
}

func LoginHandler(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputRequest request.UserRequest

		if err := c.BindJSON(&inputRequest); err != nil {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: "Bad input", Code: http.StatusBadRequest})
			return
		}
		result, token, errorToThrow := userService.Login(inputRequest)

		if errorToThrow != nil {
			utils.ThrowError(c, errorToThrow)
			return
		}
		c.Header("Authorization", "Bearer "+token)
		c.JSON(http.StatusAccepted, gin.H{
			"message": result,
		})
	}
}

func GetRoleHandler(userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//User is already authentified => authmiddleware
		userId := utils.GetUserIdInClaims(ctx)
		println("User id is ", userId)
		if userId == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "Error: no ID found in claims"})
		}
		response, err := userService.GetUserRole(userId)
		if err != nil {
			ctx.JSON(err.Code, gin.H{
				"error": err.Msg,
			})
			return
		}
		println("is admin ", response.IsAdmin)
		ctx.JSON(200, gin.H{
			"response": response,
		})
	}
}
