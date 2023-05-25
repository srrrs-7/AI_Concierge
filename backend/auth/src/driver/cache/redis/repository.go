package redis

import (
	"context"
	"errors"
	"template/util/env"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	env *env.Env
	redis *redis.Client
}

func NewRepository(env *env.Env, redis *redis.Client) *Repository {
	return &Repository{
		env: env,
		redis: redis,
	}
}

func (r *Repository) Set(ctx context.Context, key, value string) error {
	err := r.redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {
	v, err := r.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("redis key not found")
	} else if err != nil {
		return "", err
	}

	return v, nil
}