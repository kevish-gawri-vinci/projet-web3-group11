package service

import (
	entity "backend/Entity"
	request "backend/Request"
	utils "backend/Utils"
	"net/http"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll() ([]request.ArticleRequest, *utils.ErrorStruct)
	GetOneById(int) (request.ArticleRequest, *utils.ErrorStruct)
	AddArticle(request.ArticleRequest) *utils.ErrorStruct
}

type articleService struct {
	DB *gorm.DB
}

// GetAll implements ArticleService.
func (a *articleService) GetAll() ([]request.ArticleRequest, *utils.ErrorStruct) {
	db := a.DB
	var articles []entity.Article
	result := db.Find(&articles)
	if result.RowsAffected == 0 || result.Error != nil {
		return []request.ArticleRequest{}, &utils.ErrorStruct{Msg: "Erreur lors de la récupération des articles", Code: http.StatusNotFound}
	}
	var responses []request.ArticleRequest
	//Transpose the GORM entity to the request
	for i := 0; i < len(articles); i++ {
		responses = append(responses, request.ArticleRequest{
			ArticleId:   articles[i].ID,
			Name:        articles[i].Name,
			Description: articles[i].Description,
			Price:       articles[i].Price,
			ImgUrl:      articles[i].ImgUrl,
		})
	}
	return responses, nil
}

// GetOneById -> Gets an article of the ID
func (a *articleService) GetOneById(id int) (request.ArticleRequest, *utils.ErrorStruct) {
	db := a.DB
	var article entity.Article
	article.ID = id
	result := db.First(&article)

	if result.RowsAffected != 1 {
		return request.ArticleRequest{}, &utils.ErrorStruct{Msg: "Could not get the article", Code: http.StatusNotFound}
	}

	response := request.ArticleRequest{
		ArticleId:   article.ID,
		Name:        article.Name,
		Description: article.Description,
		ImgUrl:      article.ImgUrl,
		Price:       article.Price,
	}

	return response, nil
}

func (a *articleService) AddArticle(req request.ArticleRequest) *utils.ErrorStruct {
	db := a.DB
	//Create the entity for GORM
	article := entity.Article{
		ID:          req.ArticleId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	}
	result := db.Create(&article)

	if result.RowsAffected != 1 {
		return &utils.ErrorStruct{Msg: "Could not add the article", Code: http.StatusNotFound}
	}

	return nil
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{DB: db}
}
