package cache

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// redis client for sessions
func SessionsClient() *RedisStore {
	store, err := NewRedisStore(32, "tcp", os.Getenv("REDIS_ADDRS"), os.Getenv("REDIS_PWD"), []byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal("unable to connect to redis")
	}

	return &store
}

// Cliento to interact directly with redis
func Client() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRS"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	return rdb
}
