package article

import (
	"article-service/models"
	"article-service/requests"

	"github.com/gin-gonic/gin"
)

type IArticleUseCase interface {
	SearchArticles(c *gin.Context) ([]models.AlgoliaArticle, error)
	PostArticles(c *gin.Context, param requests.ParamArticle) (models.Article, error)
}
