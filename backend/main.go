package main

import (
	database "backend/Database"
	handler "backend/Handler"
	middleware "backend/Middleware"
	service "backend/Service"
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
	basketService := service.NewBasketService(db)
	orderService := service.NewOrderService(db)

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

	//Request that need authentification (eg. add article)
	/******* BASKETS **********/
	r.POST("/basket/add", middleware.AuthMiddleware(), handler.AddArticleToBasketHandler(basketService))
	r.DELETE("/basket/delete-all/:id", middleware.AuthMiddleware(), handler.DeleteBasketHandler(basketService))
	r.GET("/basket/get/:id", middleware.AuthMiddleware(), handler.GetBasketHandler(basketService))
	r.PUT("/basket/increase-quantity", middleware.AuthMiddleware(), handler.IncreaseQuantityHandler(basketService))
	r.PUT("/basket/decrease-quantity", middleware.AuthMiddleware(), handler.DecreaseQuantityHandler(basketService))

	/******* ORDERS **********/
	r.POST("/order/finalise", middleware.AuthMiddleware(), handler.FinaliseBasketHandler(orderService))
	r.GET("/order/get/:id", middleware.AuthMiddleware(), handler.GetOrderHandler(orderService))

	r.Run() // listen and serve on 0.0.0.0:8080
}
