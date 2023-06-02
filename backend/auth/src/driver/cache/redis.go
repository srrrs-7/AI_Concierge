package cache

import (
	"ai_concierge/util/env"

	"github.com/redis/rueidis"
)

func NewRedis(env *env.EnvParams[string]) (rueidis.Client, error) {
	rds, err := rueidis.NewClient(
		rueidis.ClientOption{
			InitAddress: []string{env.REDIS_URL.Value},
		},
	)
	if err != nil {
		return nil, err
	}

	return rds, nil
}
