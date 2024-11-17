package handler

import (
	request "backend/Request"
	service "backend/Service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUserHandler(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		println(c.PostForm("username"))
		println(c.PostForm("password"))
		var userRequest request.UserRequest
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data",
			})
		}
		user := userService.AddUser(userRequest)

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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Input"})
			return
		}
		result, err, token := userService.Login(inputRequest)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
		c.Header("Authorization", "Bearer "+token)
		c.JSON(http.StatusAccepted, gin.H{
			"message": result,
		})
	}
}
