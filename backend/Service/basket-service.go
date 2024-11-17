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
	AddOneArticle(request.AddOneArticleRequest) (entity.BasketItem, *utils.ErrorStruct)
	DeleteBasket(int) *utils.ErrorStruct
}

type basketService struct {
	DB *gorm.DB
}

// AddOneArticle implements BasketService.
func (b *basketService) AddOneArticle(req request.AddOneArticleRequest) (entity.BasketItem, *utils.ErrorStruct) {
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
	basketItem := entity.BasketItem{UserId: id}
	result := db.Find(&basketItem)
	if result.RowsAffected == 0 {
		return &utils.ErrorStruct{Msg: "User does not have a basket or invalid user id", Code: http.StatusBadRequest}
	}
	if result.Error != nil {
		return &utils.ErrorStruct{Msg: result.Error.Error(), Code: http.StatusInternalServerError}
	}
	return nil
}

func NewBasketService(db *gorm.DB) BasketService {
	return &basketService{DB: db}
}
