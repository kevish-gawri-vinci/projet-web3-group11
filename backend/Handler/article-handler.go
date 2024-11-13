package handler

import (
	service "backend/Service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := articleService.GetAll()
		fmt.Println(response[0])
		ctx.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}

func GetOneByIdHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawId := ctx.Params.ByName("id")
		id, err := strconv.Atoi(rawId)
		print("The converted int ", id)
		if err != nil || id == 0 {
			fmt.Errorf("Erreur dans la conversion du paramètre id")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
			})
			return
		}

		article := articleService.GetOneById(id)
		if article.ID == 0 {
			fmt.Errorf("Aucun article trouvé")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Aucun article trouvé",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": article,
		})
	}
}
