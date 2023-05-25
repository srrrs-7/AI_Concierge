package cache

import (
	"template/util/env"

	"github.com/redis/go-redis/v9"
)

func NewRedis(env *env.Env) *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redis
}
