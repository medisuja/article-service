package cache

import (
	"article-service/config"

	"github.com/go-redis/redis"
)

var (
	redisConfig = config.Config.REDIS
	redisClient *redis.Client
)

func init() {
	setupRedisConn()
}

func setupRedisConn() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.RedisHost + ":" + redisConfig.RedisPort,
		Password: redisConfig.RedisPassword,
		DB:       0,
	})
}

func RedisClient() *redis.Client {
	return redisClient
}
