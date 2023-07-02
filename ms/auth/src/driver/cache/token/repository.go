package token

import (
	"auth/util/env"
	"context"

	"github.com/redis/rueidis"
)

type RdsRepository interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type Repository struct {
	env   *env.EnvParams[string]
	redis rueidis.Client
}

func New(env *env.EnvParams[string], redis rueidis.Client) *Repository {
	return &Repository{
		env:   env,
		redis: redis,
	}
}

func (r *Repository) Set(ctx context.Context, key, value string) error {

	err := r.redis.Do(ctx, r.redis.B().Set().Key("key").Value("val").Nx().Build()).Error()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context, key string) (string, error) {

	strMap, err := r.redis.Do(ctx, r.redis.B().Get().Key(key).Build()).AsStrMap()
	if err != nil {
		return "", err
	}

	return strMap[key], nil
}
