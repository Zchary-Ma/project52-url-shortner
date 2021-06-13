package storage

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zchary-ma/url-shortener/config"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Exists(key string) (bool, error)
}

type RedisStorage struct {
	client *redis.Client
}

func CreateClient(c config.Configurations) *RedisStorage {
	var storage = &RedisStorage{}
	storage.client = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	_, err := storage.client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connecting redis server succeeds.")
	return storage
}

func (r *RedisStorage) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RedisStorage) Set(key string, value interface{}) error {
	return r.client.Set(key, value, 0).Err()
}

func (r *RedisStorage) Exists(key string) (bool, error) {
	ok, err := r.client.Exists(key).Result()
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}

func (r RedisStorage) String() string {
	return "RedisStorage"
}
