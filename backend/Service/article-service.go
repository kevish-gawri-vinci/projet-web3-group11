package service

import (
	entity "backend/Entity"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll() []entity.Article
}

type articleService struct {
	DB *gorm.DB
}

// GetAll implements ArticleService.
func (a *articleService) GetAll() []entity.Article {
	db := a.DB
	var articles []entity.Article
	db.Find(&articles)
	return articles
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{DB: db}
}
