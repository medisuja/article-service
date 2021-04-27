package usecase

import (
	articleInterfaces "article-service/app/article"
	redisInterfaces "article-service/app/redis"
	"article-service/config"
	"article-service/helpers"
	"article-service/models"
	"article-service/requests"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

type articleUsecase struct {
	articleRepo articleInterfaces.IArticleRepository
	redisRepo   redisInterfaces.IRedisRepository
}

func NewArticleUsecase(art articleInterfaces.IArticleRepository, rds redisInterfaces.IRedisRepository) articleInterfaces.IArticleUseCase {
	return &articleUsecase{
		articleRepo: art,
		redisRepo:   rds,
	}
}

func (a *articleUsecase) SearchArticles(c *gin.Context) ([]models.AlgoliaArticle, error) {
	var (
		articlesData []models.AlgoliaArticle
		articles     algoliasearch.QueryRes
	)

	keyword := c.Query("query")
	if keyword != "" {
		keywordSplit := strings.Split(keyword, "")

		if len(keywordSplit) < 3 {
			return nil, errors.New("Keyword must have greater than 3 character")
		}
	}

	limit := config.Config.Algolia.HitsPerPage
	articles, err := helpers.SearchAlgolia(config.Config.Algolia.IndicesArticle, keyword, limit)
	for _, article := range articles.Hits {
		articlesData = append(articlesData, assignArticleAlgolia(article))
	}

	if err != nil {
		return articlesData, err
	}

	return articlesData, nil
}

func assignArticleAlgolia(articleData algoliasearch.Map) models.AlgoliaArticle {
	fmt.Println(int(articleData["created"].(float64)))

	var result models.AlgoliaArticle
	objectID, _ := strconv.Atoi(articleData["objectID"].(string))
	result.ObjectID = objectID
	result.Author = articleData["author"].(string)
	result.Title = articleData["title"].(string)
	result.Slug = articleData["slug"].(string)
	result.Body = articleData["body"].(string)
	result.CreatedAt = int64(articleData["created"].(float64))

	return result
}

func (a *articleUsecase) PostArticles(c *gin.Context, request requests.ParamArticle) (models.Article, error) {
	var article models.Article

	article.Author = request.Author
	article.Title = request.Title
	article.Slug = slug.Make(article.Title)
	article.Body = request.Body

	creatArticle, err := a.articleRepo.Create(article)
	if err != nil {
		return creatArticle, err
	}

	// stor to algolia
	articles := []interface{}{assignAlgoliaValue(&creatArticle)}
	_, errorStore := helpers.AddBatchObject(config.Config.Algolia.IndicesArticle, articles)
	if errorStore != nil {
		return article, errorStore
	}

	// store to redis (cache)
	key := "article:" + creatArticle.Slug
	json, _ := json.Marshal(creatArticle)
	go a.redisRepo.Set(key, string(json), time.Minute*time.Duration(config.Config.REDIS.ExpiredMinute))

	return creatArticle, nil
}

func assignAlgoliaValue(article *models.Article) *models.AlgoliaArticle {
	algoliaArticle := new(models.AlgoliaArticle)
	algoliaArticle.ObjectID = article.ID
	algoliaArticle.Author = article.Author
	algoliaArticle.Title = article.Title
	algoliaArticle.Body = article.Body
	algoliaArticle.Slug = article.Slug
	algoliaArticle.CreatedAt = article.CreatedAt.Unix()

	return algoliaArticle
}
