package service

import (
	entity "backend/Entity"
	"errors"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll() ([]entity.Article, error)
	GetOneById(int) (entity.Article, error)
}

type articleService struct {
	DB *gorm.DB
}

// GetAll implements ArticleService.
func (a *articleService) GetAll() ([]entity.Article, error) {
	db := a.DB
	var articles []entity.Article
	result := db.Find(&articles)
	if result.RowsAffected == 0 || result.Error != nil {
		return articles, errors.New("Erreur lors de la récupération des articles")
	}
	return articles, nil
}

// GetOneById -> Gets an article of the ID
func (a *articleService) GetOneById(id int) (entity.Article, error) {
	db := a.DB
	var article entity.Article
	article.ID = id
	result := db.First(&article)

	if result.RowsAffected != 1 {
		return article, errors.New("Error in getting the article")
	}

	return article, nil
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{DB: db}
}
