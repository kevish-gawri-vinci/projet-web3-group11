package handler

import (
	service "backend/Service"
	utils "backend/Utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FinaliseBasketHandler(orderService service.OrderService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//retrieve ID from ctx variables
		userId := utils.GetUserIdInClaims(ctx)

		order, err := orderService.FinaliseBasket(userId)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response": order,
		})
	}
}

func GetOrderHandler(orderService service.OrderService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Get id from query
		rawId := ctx.Params.ByName("id")
		id, err := strconv.Atoi(rawId)
		if err != nil || id == 0 {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: err.Error(), Code: http.StatusBadRequest})
			return
		}
		userId := utils.GetUserIdInClaims(ctx)

		fullOrder, errorToThrow := orderService.GetOrder(id, userId)

		if errorToThrow != nil {
			utils.ThrowError(ctx, errorToThrow)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"response": fullOrder,
		})
	}
}

func GetAllOrdersHandler(orderService service.OrderService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := utils.GetUserIdInClaims(ctx)
		orderList, err := orderService.GetAllOrders(userId)
		if err != nil {
			utils.ThrowError(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"response": orderList,
		})
		return
	}
}
