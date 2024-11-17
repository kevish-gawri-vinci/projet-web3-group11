package handler

import (
	request "backend/Request"
	service "backend/Service"
	utils "backend/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddArticleToBasketHandler(basketService service.BasketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.AddOneArticleRequest
		ctx.BindJSON(&req)
		basketItem, err := basketService.AddOneArticle(req)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"basket_item": basketItem,
		})
	}
}
