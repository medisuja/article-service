package app

import (
	ArticleInterfaces "article-service/app/article"
	ArticleHandler "article-service/app/article/handler"
	HealthCheckInterfaces "article-service/app/health-check"
	HealthCheckHandler "article-service/app/health-check/handler"

	"github.com/gin-gonic/gin"
)

// Health Check
func HealthCheckHttpHandler(r *gin.Engine, useCase HealthCheckInterfaces.IHealthCheckUsecase) {
	handler := &HealthCheckHandler.HealthCheckHandler{HealthCheckUsecase: useCase}
	r.GET("/", handler.Check)
}

// Get Article
func ArticleHttpHandler(r *gin.Engine, useCase ArticleInterfaces.IArticleUseCase) {
	handler := &ArticleHandler.ArticleHandler{Usecase: useCase}
	r.POST("/articles", handler.PostArticles)
	r.GET("/articles", handler.SearchArticles)
}
