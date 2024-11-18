package main

import (
	database "backend/Database"
	handler "backend/Handler"
	service "backend/Service"
	middleware "backend/Middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db, err := database.InitDb()

	if err != nil {
		fmt.Errorf("A problem occured when connecting to the database")
	}

	userService := service.NewUserService(db)
	articleService := service.NewArticleService(db)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	r.GET("/ping", func(c *gin.Context) {
		println(c.Query("id"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/adduser", handler.AddUserHandler(userService))
	r.GET("/article/getall", handler.GetAllHandler(articleService))
	r.GET("/article/get/:id", handler.GetOneByIdHandler(articleService))
	r.POST("/auth/login", handler.LoginHandler(userService))

	r.Run() // listen and serve on 0.0.0.0:8080
}
