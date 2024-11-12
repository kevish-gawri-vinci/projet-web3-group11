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
		var userRequest request.AddUserRequest
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data",
			})
		}
		user := userService.AddUser(userRequest)

		c.JSON(http.StatusCreated, gin.H{
			"id":       user.Id,
			"username": user.Username,
		})
	}
}
