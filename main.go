package main

import (
	"fmt"
	"log"

	routes "article-service/app"
	ArticleRepo "article-service/app/article/repository"
	ArticleUsecase "article-service/app/article/usecase"
	HealthCheckRepo "article-service/app/health-check/repository"
	HealthCheckUsecase "article-service/app/health-check/usecase"
	RedisRepo "article-service/app/redis/repository"
	"article-service/cache"
	"article-service/config"
	gorm "article-service/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var appConfig = config.Config.App

func init() {

}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
	}))

	// setting database connection
	dbConnection := gorm.MysqlConn()

	redis := cache.RedisClient()
	redisRepo := RedisRepo.NewRedisRepository(redis)

	// Routes Health Check
	healthCheckRepo := HealthCheckRepo.NewHealthCheckRepository(dbConnection)
	halthCheckUsecase := HealthCheckUsecase.NewHealthCheckUsecase(healthCheckRepo)
	routes.HealthCheckHttpHandler(r, halthCheckUsecase)

	// Routes Articles
	articleRepo := ArticleRepo.NewArticleRepository(dbConnection)
	articleUscase := ArticleUsecase.NewArticleUsecase(articleRepo, redisRepo)
	routes.ArticleHttpHandler(r, articleUscase)

	// Server
	if err := r.Run(fmt.Sprintf(":%s", appConfig.HttpPort)); err != nil {
		log.Fatal(err)
	}
}
