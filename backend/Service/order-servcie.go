package service

import (
	entity "backend/Entity"
	request "backend/Request"
	utils "backend/Utils"
	"net/http"

	"gorm.io/gorm"
)

type OrderService interface {
	FinaliseBasket(int) *utils.ErrorStruct
	GetAllOrders(int) ([]request.OrderListLine, *utils.ErrorStruct)
	GetOrder(int, int) (request.FullOrder, *utils.ErrorStruct)
}

type orderService struct {
	DB            *gorm.DB
	BasketService BasketService
}

// FinaliseBasket implements OrderService. //TODO
func (o *orderService) FinaliseBasket(userId int) *utils.ErrorStruct {
	db := o.DB

	//Get the basket of the user
	basket, errorToThrow := o.BasketService.GetBasket(userId)
	if errorToThrow != nil {
		return errorToThrow
	}
	// Add the order in the DB
	order := entity.Order{UserId: userId}
	result := db.Create(&order)
	if result.Error != nil {
		errorToThrow := &utils.ErrorStruct{Msg: "Error while creating the order", Code: http.StatusInternalServerError}
		return errorToThrow
	}

	//For each line of the basket, create an orderLine
	articles := basket.Articles
	for i := 0; i < len(articles); i++ {
		orderLine := entity.OrderLine{OrderId: order.ID, Quantity: articles[i].Quantity, ArticleId: articles[i].ArticleDetail.ID}
		result := db.Create(&orderLine)
		if result.Error != nil {
			errorToThrow := &utils.ErrorStruct{Msg: "Error while adding the article in order", Code: http.StatusInternalServerError}
			return errorToThrow
		}
	}

	//Delete Basket
	errorToThrow = o.BasketService.DeleteBasket(userId)
	if errorToThrow != nil {
		return errorToThrow
	}

	return nil
}

// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders(userId int) ([]request.OrderListLine, *utils.ErrorStruct) {
	db := o.DB
	var orderList []request.OrderListLine
	var allUserOrders []entity.Order
	result := db.Where("user_id = ?", userId).Find(&allUserOrders)
	if result.Error != nil {
		return []request.OrderListLine{}, &utils.ErrorStruct{Msg: "Error in the database", Code: http.StatusInternalServerError}
	}
	if len(allUserOrders) == 0 {
		return []request.OrderListLine{}, &utils.ErrorStruct{Msg: "No order found", Code: http.StatusNotFound}
	}
	for i := 0; i < len(allUserOrders); i++ {
		//For each order, find its order lines to calculate order total price and total quantity
		var orderLines []entity.OrderLine //All lines of order where order_id = allUserOrders[i].ID
		result := db.Where("order_id = ?", allUserOrders[i].ID).Find(&orderLines)
		if result.Error != nil {
			return []request.OrderListLine{}, &utils.ErrorStruct{Msg: "Error in the database", Code: http.StatusInternalServerError}
		}
		var totalQuantity int
		var totalPrice float32
		if len(orderLines) == 0 { // If the order has no line (unlikely)
			continue
		}
		for j := 0; j < len(orderLines); j++ {
			article := entity.Article{ID: orderLines[j].ArticleId}
			totalQuantity += orderLines[j].Quantity
			result := db.Find(&article)
			if result.Error != nil {
				return []request.OrderListLine{}, &utils.ErrorStruct{Msg: "Error in the database", Code: http.StatusInternalServerError}
			}
			totalPrice += (float32(orderLines[j].Quantity) * float32(article.Price))
		}
		//Set the order list line and append to the orderlist slice
		orderListLine := request.OrderListLine{
			OrderId:       allUserOrders[i].ID,
			TotalPrice:    totalPrice,
			TotalQuantity: totalQuantity,
		}
		orderList = append(orderList, orderListLine)
	}
	return orderList, nil
}

// GetOrder implements OrderService.
func (o *orderService) GetOrder(orderId int, userId int) (request.FullOrder, *utils.ErrorStruct) {
	db := o.DB
	//Get every order line of that orderId
	var orderLines []entity.OrderLine
	result := db.Where("order_id = ?", orderId).Find(&orderLines)
	if result.Error != nil {
		errrorToThrow := &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
		return request.FullOrder{}, errrorToThrow
	}
	if result.RowsAffected == 0 {
		errrorToThrow := &utils.ErrorStruct{Msg: "No order found", Code: http.StatusNotFound}
		return request.FullOrder{}, errrorToThrow
	}

	//Get the articles of that order
	var orderDetails []request.OrderDetail
	var totalOrderPrice float32 = 0.0
	for i := 0; i < len(orderLines); i++ {
		//Check if the order is the order of the userID
		order := entity.Order{ID: orderLines[i].OrderId}
		result := db.Find(&order)
		if result.Error != nil {
			errrorToThrow := &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
			return request.FullOrder{}, errrorToThrow
		}
		if order.UserId != userId {
			return request.FullOrder{}, &utils.ErrorStruct{Msg: "Unauthorized", Code: http.StatusUnauthorized}
		}
		var orderDetail request.OrderDetail
		var article entity.Article = entity.Article{ID: orderLines[i].ArticleId}
		var articleLine request.ArticleLine = request.ArticleLine{ArticleId: orderLines[i].ArticleId, Quantity: orderLines[i].Quantity}
		result = db.Find(&article)
		if result.Error != nil {
			errrorToThrow := &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
			return request.FullOrder{}, errrorToThrow
		}
		//Calculate price of the order
		var orderLinePrice float32
		orderLinePrice = article.Price * float32(orderLines[i].Quantity)
		articleLine.Price = orderLinePrice
		//Set the orderDetail
		orderDetail.ArticleDetail = article
		orderDetail.ArticleLine = articleLine
		// Add the article to the array (slice)
		orderDetails = append(orderDetails, orderDetail)
		//Calculate total order price
		totalOrderPrice += orderLinePrice
	}

	//Set the return FullOrder
	var fullOrder request.FullOrder
	fullOrder.Articles = orderDetails
	fullOrder.OrderId = orderId
	fullOrder.TotalPrice = totalOrderPrice
	return fullOrder, nil
}

func NewOrderService(db *gorm.DB) OrderService {
	basketService := NewBasketService(db)
	return &orderService{DB: db, BasketService: basketService}
}
