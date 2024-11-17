package handler

import (
	service "backend/Service"
	utils "backend/Utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := articleService.GetAll()

		if err != nil {
			// Send error to the middleware
			println("test erreur ")
			utils.ThrowError(ctx, err)
		}
		println("blabla")
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
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: err.Error(), Code: http.StatusBadRequest})
			return
		}

		article, error := articleService.GetOneById(id)
		if err != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": article,
		})
	}
}
