package main

import (
	database "backend/Database"
	handler "backend/Handler"
	middleware "backend/Middleware"
	service "backend/Service"
	"fmt"

	"github.com/gin-contrib/cors"
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

	//Cors option
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // URL du frontend React
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Méthodes HTTP autorisées
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // En-têtes autorisés
		ExposeHeaders:    []string{"Authorization"},                           // En-têtes à exposer pour la lecture côté client
		AllowCredentials: true,                                                // Autoriser les cookies et les informations d'identification
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*************** USER **************/
	userGroup := r.Group("/auth")
	userGroup.POST("/signup", handler.AddUserHandler(userService))
	userGroup.POST("/login", handler.LoginHandler(userService))
	userGroup.GET("/user-role", middleware.AuthMiddleware(), handler.GetRoleHandler(userService)) //Auth

	/****** ARTICLES *********/
	articleGroup := r.Group("/article")
	articleGroup.GET("/getall", handler.GetAllHandler(articleService))
	articleGroup.GET("/get/:id", handler.GetOneByIdHandler(articleService))
	articleGroup.POST("/add", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.AddArticleHandler(articleService)) //Auth + admin

	//Request that need authentification (eg. add article)
	/******* BASKETS **********/
	basketGroup := r.Group("/basket")
	basketGroup.POST("/add", middleware.AuthMiddleware(), handler.AddArticleToBasketHandler(basketService))
	basketGroup.DELETE("/delete-all", middleware.AuthMiddleware(), handler.DeleteBasketHandler(basketService))
	basketGroup.GET("/get", middleware.AuthMiddleware(), handler.GetBasketHandler(basketService))
	basketGroup.PUT("/increase-quantity", middleware.AuthMiddleware(), handler.IncreaseQuantityHandler(basketService))
	basketGroup.PUT("/decrease-quantity", middleware.AuthMiddleware(), handler.DecreaseQuantityHandler(basketService))

	/******* ORDERS **********/
	orderGroup := r.Group("/order")
	orderGroup.POST("/finalize", middleware.AuthMiddleware(), handler.FinaliseBasketHandler(orderService))
	orderGroup.GET("/get/:id", middleware.AuthMiddleware(), handler.GetOrderHandler(orderService))
	orderGroup.GET("/getall", middleware.AuthMiddleware(), handler.GetAllOrdersHandler(orderService))

	r.Run() // listen and serve on 0.0.0.0:8080
}
