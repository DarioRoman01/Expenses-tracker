package cache

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
)

type RedisStore interface {
	Store
}

func NewRedisStore(size int, network, address, password string, keyPairs ...[]byte) (RedisStore, error) {
	store, err := redistore.NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &redisStore{store}, nil
}

type redisStore struct {
	*redistore.RediStore
}

func (c *redisStore) Options(options Options) {
	c.RediStore.Options = &sessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}

func (c *redisStore) MaxAge(age int) {
	c.RediStore.SetMaxAge(age)
}

// Default: 4096,
func (c *redisStore) MaxLength(length int) {
	c.RediStore.SetMaxLength(length)
}
