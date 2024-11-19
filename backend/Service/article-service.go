package service

import (
	entity "backend/Entity"
	utils "backend/Utils"
	"net/http"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll() ([]entity.Article, *utils.ErrorStruct)
	GetOneById(int) (entity.Article, *utils.ErrorStruct)
	AddArticle(req entity.Article) *utils.ErrorStruct
}

type articleService struct {
	DB *gorm.DB
}

// GetAll implements ArticleService.
func (a *articleService) GetAll() ([]entity.Article, *utils.ErrorStruct) {
	db := a.DB
	var articles []entity.Article
	result := db.Find(&articles)
	if result.RowsAffected == 0 || result.Error != nil {
		return articles, &utils.ErrorStruct{Msg: "Erreur lors de la récupération des articles", Code: http.StatusNotFound}
	}
	return articles, nil
}

// GetOneById -> Gets an article of the ID
func (a *articleService) GetOneById(id int) (entity.Article, *utils.ErrorStruct) {
	db := a.DB
	var article entity.Article
	article.ID = id
	result := db.First(&article)

	if result.RowsAffected != 1 {
		return article, &utils.ErrorStruct{Msg: "Could not get the article", Code: http.StatusNotFound}
	}

	return article, nil
}

func (a *articleService) AddArticle(req entity.Article) *utils.ErrorStruct {
	db := a.DB
	result := db.Create(&req)

	if result.RowsAffected != 1 {
		return &utils.ErrorStruct{Msg: "Could not add the article", Code: http.StatusNotFound}
	}

	return nil
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{DB: db}
}
