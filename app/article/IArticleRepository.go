package article

import "article-service/models"

type IArticleRepository interface {
	Fetch() ([]*models.Article, error)
	Create(models.Article) (models.Article, error)
}
