package cache

import (
	"auth/util/env"

	"github.com/redis/rueidis"
)

func New(env *env.EnvParams[string]) (rueidis.Client, error) {
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
