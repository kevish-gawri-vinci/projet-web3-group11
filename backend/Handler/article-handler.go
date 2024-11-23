package handler

import (
	request "backend/Request"
	service "backend/Service"
	utils "backend/Utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := articleService.GetAll()

		if err != nil {
			// Send error to the middleware
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": response,
		})
		return
	}
}

func GetOneByIdHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawId := ctx.Params.ByName("id")
		id, err := strconv.Atoi(rawId)
		if err != nil || id == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: err.Error(), Code: http.StatusBadRequest})
			return
		}

		article, error := articleService.GetOneById(id)
		if error != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": article,
		})
		return
	}
}

func AddArticleHandler(articleService service.ArticleService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Check if user is admin
		//Request article is the same as entity article to go in gorm so I will take the same
		var req request.ArticleRequest
		errorInBind := ctx.ShouldBind(&req)
		if errorInBind != nil {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "Bad input", Code: http.StatusBadRequest})
			return
		}
		err := articleService.AddArticle(req)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{})
		return
	}
}
