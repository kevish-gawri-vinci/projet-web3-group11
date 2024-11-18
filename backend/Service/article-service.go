package service

import (
	entity "backend/Entity"

	"gorm.io/gorm"
)

type ArticleService interface {
	GetAll() []entity.Article
	GetOneById(int) entity.Article
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

// GetOneById -> Gets an article of the ID
func (a *articleService) GetOneById(id int) entity.Article {
	db := a.DB
	var article entity.Article
	article.ID = id
	result := db.First(&article)

	if result.RowsAffected != 1 {
		return entity.Article{ID: 0}
	}

	return article
}

func NewArticleService(db *gorm.DB) ArticleService {
	return &articleService{DB: db}
}
