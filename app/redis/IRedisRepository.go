package redis

import "time"

type IRedisRepository interface {
	Set(key string, value interface{}, ExpireAt time.Duration) error
	Get(key string) ([]byte, error)
	Delete(keys []string) error
	Publish(channel string, messages string)
}
