package service

import (
	entity "backend/Entity"
	request "backend/Request"
	utils "backend/Utils"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type BasketService interface {
	AddOneArticle(request.BasketArticleRequest) (entity.BasketItem, *utils.ErrorStruct)
	DeleteBasket(int) *utils.ErrorStruct
	GetBasket(int) (entity.Basket, *utils.ErrorStruct)
	IncreaseQuantity(request.BasketArticleRequest) *utils.ErrorStruct
	DecreaseQuantity(request.BasketArticleRequest) *utils.ErrorStruct
}

type basketService struct {
	DB *gorm.DB
}

// AddOneArticle implements BasketService.
func (b *basketService) AddOneArticle(req request.BasketArticleRequest) (entity.BasketItem, *utils.ErrorStruct) {
	//SQL query
	db := b.DB
	//If user hasn't added the product in the basket yet
	basketItem := entity.BasketItem{ArticleId: req.ArticleId, UserId: req.UserId, Quantity: req.Quantity}
	result := db.Create(&basketItem)

	if result.Error != nil {
		//If the article is already in the basket of the user
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				// Get the quantity of the row of article ID and user ID
				basketItem := entity.BasketItem{ArticleId: req.ArticleId, UserId: req.UserId}
				db.First(&basketItem)
				newQuantity := req.Quantity + basketItem.Quantity
				db.Model(&basketItem).Where("user_id = ? AND article_id = ?", req.UserId, req.ArticleId).Update("quantity", newQuantity)
				return basketItem, nil
			} else {
				return entity.BasketItem{}, &utils.ErrorStruct{Msg: "Could not add article to the basket", Code: http.StatusInternalServerError}
			}
		}
	}
	return basketItem, nil
}

func (b *basketService) DeleteBasket(id int) *utils.ErrorStruct {
	db := b.DB
	basketItems := []entity.BasketItem{}
	result := db.Where("user_id = ?", id).Delete(&basketItems)

	if result.Error != nil {
		return &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
	}

	if result.RowsAffected == 0 {
		return &utils.ErrorStruct{Msg: "User does not have a basket or invalid user id", Code: http.StatusBadRequest}
	}

	return nil
}

func (b *basketService) GetBasket(id int) (entity.Basket, *utils.ErrorStruct) {
	db := b.DB
	basketItems := []entity.BasketItem{}
	result := db.Where("user_id = ?", id).Find(&basketItems)
	println("In the service of GetBasket")
	if result.Error != nil {
		return entity.Basket{}, &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
	}

	if result.RowsAffected == 0 {
		return entity.Basket{}, &utils.ErrorStruct{Msg: "User does not have a basket or invalid user id", Code: http.StatusBadRequest}
	}
	var totalPrice float32
	println("Size of articles ", len(basketItems))
	for i := 0; i < len(basketItems); i++ {
		article := entity.Article{ID: basketItems[i].ArticleId}
		db.Find(&article)
		totalPrice += float32(article.Price * float32(basketItems[i].Quantity))
	}
	basket := entity.Basket{
		Articles:   basketItems,
		TotalPrice: float32(totalPrice),
		UserId:     id,
	}
	return basket, nil
}

func (b *basketService) IncreaseQuantity(req request.BasketArticleRequest) *utils.ErrorStruct {
	db := b.DB
	basketItem := entity.BasketItem{ArticleId: req.ArticleId, UserId: req.UserId}
	result1 := db.Find(&basketItem)
	if result1.RowsAffected != 1 {
		return &utils.ErrorStruct{Msg: "Article not found", Code: http.StatusNotFound}
	}
	if result1.Error != nil {
		return &utils.ErrorStruct{Msg: "Something went wrong in the database", Code: http.StatusInternalServerError}
	}
	println(basketItem.Quantity)
	newQuantity := basketItem.Quantity + req.Quantity
	result2 := db.Model(&basketItem).Update("quantity", newQuantity)

	if result2.Error != nil {
		return &utils.ErrorStruct{Msg: "Something went wrong in the database", Code: http.StatusInternalServerError}
	}

	return nil
}

func (b *basketService) DecreaseQuantity(req request.BasketArticleRequest) *utils.ErrorStruct {
	db := b.DB
	basketItem := entity.BasketItem{ArticleId: req.ArticleId, UserId: req.UserId}
	result1 := db.Find(&basketItem)
	if result1.RowsAffected != 1 {
		return &utils.ErrorStruct{Msg: "Article not found", Code: http.StatusNotFound}
	}
	if result1.Error != nil {
		return &utils.ErrorStruct{Msg: "Something went wrong in the database", Code: http.StatusInternalServerError}
	}
	if basketItem.Quantity < req.Quantity {
		return &utils.ErrorStruct{Msg: "Invalid quantity", Code: http.StatusBadRequest}
	}
	if basketItem.Quantity == req.Quantity {
		db.Delete(&basketItem)
		return nil
	}
	newQuantity := basketItem.Quantity - req.Quantity
	result2 := db.Model(&basketItem).Update("quantity", newQuantity)

	if result2.Error != nil {
		return &utils.ErrorStruct{Msg: "Something went wrong in the database", Code: http.StatusInternalServerError}
	}
	return nil
}

func NewBasketService(db *gorm.DB) BasketService {
	return &basketService{DB: db}
}
