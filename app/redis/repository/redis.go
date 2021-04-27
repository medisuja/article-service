package repository

import (
	"time"

	RedisInterface "article-service/app/redis"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(Client *redis.Client) RedisInterface.IRedisRepository {
	return &RedisRepository{Client}
}

func (m *RedisRepository) Set(key string, value interface{}, ExpireAt time.Duration) error {
	err := m.Client.Set(key, value, ExpireAt).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *RedisRepository) Get(key string) ([]byte, error) {
	resp, err := m.Client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *RedisRepository) Delete(keys []string) error {
	err := m.Client.Del(keys...).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *RedisRepository) Publish(channel string, messages string) {
	m.Client.Publish(channel, messages)
	return
}
