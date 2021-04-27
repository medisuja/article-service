package handler

import (
	articleInterfaces "article-service/app/article"
	"article-service/helpers"
	"article-service/requests"
	"article-service/transformers"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type paramArticleRules struct {
	Author string `valid:"required~parameter is empty"`
	Title  string `valid:"required~parameter is empty"`
	Body   string `valid:"required~parameter is empty"`
}

type ArticleHandler struct {
	Usecase articleInterfaces.IArticleUseCase
}

func (a *ArticleHandler) SearchArticles(c *gin.Context) {
	articles, err := a.Usecase.SearchArticles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(articles) < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	// Transform
	res := new(transformers.CollectionTransformer)
	res.TransformArticleAlgoliaCollection(articles)

	c.JSON(http.StatusOK, res)
	return
}

func (a *ArticleHandler) PostArticles(c *gin.Context) {
	request := requests.ParamArticle{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validation := ValidatePostArticle(request)
	if validation != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validation})
		return
	}

	postArticle, errorPost := a.Usecase.PostArticles(c, request)
	if errorPost != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorPost.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postArticle})
	return
}

func ValidatePostArticle(params requests.ParamArticle) interface{} {
	authorRules := &paramArticleRules{
		Author: params.Author,
		Title:  params.Title,
		Body:   params.Body,
	}

	_, err := govalidator.ValidateStruct(authorRules)
	if err != nil {
		respErr := helpers.ValidationError(authorRules, err)
		if respErr != nil {
			return respErr
		}
	}

	return nil
}
