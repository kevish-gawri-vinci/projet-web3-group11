package main

import (
	handler "backend/Handler"
	service "backend/Service"

	"github.com/gin-gonic/gin"
)

func main() {
	userService := service.NewUserService()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		println(c.Query("id"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/adduser", handler.AddUserHandler(userService))
	r.Run() // listen and serve on 0.0.0.0:8080
}
