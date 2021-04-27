package repository

import (
	articleInterface "article-service/app/article"
	"article-service/models"

	"github.com/jinzhu/gorm"
)

type articleRepository struct {
	Connection *gorm.DB
}

func NewArticleRepository(Connection *gorm.DB) articleInterface.IArticleRepository {
	return &articleRepository{Connection}
}

func (m *articleRepository) Fetch() ([]*models.Article, error) {
	var (
		articles []*models.Article
		err      error
	)

	// Initialization
	tx := m.Connection.Begin()
	err = tx.Find(&articles).Error
	if err != nil {
		tx.Rollback()
		return articles, err
	}

	tx.Commit()

	return articles, err
}

func (a *articleRepository) Create(data models.Article) (models.Article, error) {
	article := models.Article{
		Author: data.Author,
		Title:  data.Title,
		Slug:   data.Slug,
		Body:   data.Body,
	}

	// Initialization
	tx := a.Connection.Begin()
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return article, err
	}

	tx.Commit()

	return article, nil
}
