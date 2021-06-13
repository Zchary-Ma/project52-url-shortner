package storage

import (
	"github.com/go-redis/redis"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Exists(key string) (bool, error)
}

func CreateClient() *redis.Client {
}
