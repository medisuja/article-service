package transformers

import (
	"article-service/models"
	"time"
)

type (
	Article struct {
		Author    string    `json:"author"`
		Title     string    `json:"title"`
		Slug      string    `json:"slug"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created"`
	}
)

// Transform for search station
func (res *CollectionTransformer) TransformArticleAlgoliaCollection(articles []models.AlgoliaArticle) {
	for _, article := range articles {
		res.Data = append(res.Data, assignAlgoliaArticle(article))
	}
}

func assignAlgoliaArticle(article models.AlgoliaArticle) interface{} {
	result := Article{}
	result.Author = article.Author
	result.Title = article.Title
	result.Slug = article.Slug
	result.Body = article.Body
	result.CreatedAt = time.Unix(article.CreatedAt, 0)
	return result
}
