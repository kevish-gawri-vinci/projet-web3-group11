package handler

import (
	service "backend/Service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		print("Dans le handler")
		response := articleService.GetAll()
		fmt.Println(response[0])
		ctx.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}
