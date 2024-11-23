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
		var req request.BasketArticleRequest
		ctx.BindJSON(&req)
		if req.Quantity == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "Quantity cannot be 0", Code: http.StatusBadRequest})
			return
		}
		req.UserId = utils.GetUserIdInClaims(ctx)
		err := basketService.AddOneArticle(req)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func DeleteBasketHandler(basketService service.BasketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := utils.GetUserIdInClaims(ctx)
		error := basketService.DeleteBasket(id)
		if error != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func GetBasketHandler(basketService service.BasketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := utils.GetUserIdInClaims(ctx)
		basket, error := basketService.GetBasket(id)

		if error != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": basket,
		})
	}
}

func IncreaseQuantityHandler(basketService service.BasketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.BasketArticleRequest
		ctx.BindJSON(&req)
		userId := utils.GetUserIdInClaims(ctx)
		req.UserId = userId
		error := basketService.IncreaseQuantity(req)
		if error != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func DecreaseQuantityHandler(basketService service.BasketService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request.BasketArticleRequest
		ctx.BindJSON(&req)
		userId := utils.GetUserIdInClaims(ctx)
		req.UserId = userId
		//If Quantity is 0
		if req.Quantity == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "Invalid quantity", Code: http.StatusBadRequest})
			return
		}

		error := basketService.DecreaseQuantity(req)
		if error != nil {
			utils.ThrowError(ctx, error)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
