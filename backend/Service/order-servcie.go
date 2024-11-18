package service

import (
	entity "backend/Entity"
	utils "backend/Utils"
	"net/http"

	"gorm.io/gorm"
)

type OrderService interface {
	FinaliseBasket(int) (entity.Order, *utils.ErrorStruct)
	GetAllOrders(int) ([]entity.Order, *utils.ErrorStruct)
	GetOrder(int) (entity.FullOrder, *utils.ErrorStruct)
}

type orderService struct {
	DB            *gorm.DB
	BasketService BasketService
}

// FinaliseBasket implements OrderService. //TODO
func (o *orderService) FinaliseBasket(userId int) (entity.Order, *utils.ErrorStruct) {
	db := o.DB

	//Get the basket of the user
	basket, errorToThrow := o.BasketService.GetBasket(userId)
	if errorToThrow != nil {
		return entity.Order{}, errorToThrow
	}
	// Add the order in the DB
	order := entity.Order{UserId: userId}
	result := db.Create(&order)
	if result.Error != nil {
		errorToThrow := &utils.ErrorStruct{Msg: "Error while creating the order", Code: http.StatusInternalServerError}
		return entity.Order{}, errorToThrow
	}

	//For each line of the basket, create an orderLine
	articles := basket.Articles
	for i := 0; i < len(articles); i++ {
		orderLine := entity.OrderLine{OrderId: order.ID, Quantity: articles[i].Quantity, ArticleId: articles[i].ArticleId}
		result := db.Create(&orderLine)
		if result.Error != nil {
			errorToThrow := &utils.ErrorStruct{Msg: "Error while adding the article in order", Code: http.StatusInternalServerError}
			return entity.Order{}, errorToThrow
		}
	}

	//Delete Basket
	errorToThrow = o.BasketService.DeleteBasket(userId)
	if errorToThrow != nil {
		return entity.Order{}, errorToThrow
	}

	return order, nil
}

// TODO
// GetAllOrders implements OrderService.
func (o *orderService) GetAllOrders(userId int) ([]entity.Order, *utils.ErrorStruct) {
	panic("unimplemented")
}

// GetOrder implements OrderService.
func (o *orderService) GetOrder(orderId int) (entity.FullOrder, *utils.ErrorStruct) {
	db := o.DB
	//Get every order line of that orderId
	var orderLines []entity.OrderLine
	result := db.Where("order_id = ?", orderId).Find(&orderLines)
	if result.Error != nil {
		errrorToThrow := &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
		return entity.FullOrder{}, errrorToThrow
	}
	if result.RowsAffected == 0 {
		errrorToThrow := &utils.ErrorStruct{Msg: "No order found", Code: http.StatusNotFound}
		return entity.FullOrder{}, errrorToThrow
	}

	//Get the articles of that order
	var orderDetails []entity.OrderDetail
	var totalOrderPrice float32 = 0.0
	for i := 0; i < len(orderLines); i++ {
		var orderDetail entity.OrderDetail
		var article entity.Article = entity.Article{ID: orderLines[i].ArticleId}
		var articleLine entity.ArticleLine = entity.ArticleLine{ArticleId: orderLines[i].ArticleId, Quantity: orderLines[i].Quantity}
		result := db.Find(&article)
		if result.Error != nil {
			errrorToThrow := &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
			return entity.FullOrder{}, errrorToThrow
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
	var fullOrder entity.FullOrder
	fullOrder.Articles = orderDetails
	fullOrder.OrderId = orderId
	fullOrder.TotalPrice = totalOrderPrice
	return fullOrder, nil
}

func NewOrderService(db *gorm.DB) OrderService {
	basketService := NewBasketService(db)
	return &orderService{DB: db, BasketService: basketService}
}
