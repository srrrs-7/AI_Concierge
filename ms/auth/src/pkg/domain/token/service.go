package token

import (
	"auth/driver/cache/token"
	"auth/util/env"
)

type UseCase interface {
}

type Service struct {
	env   *env.EnvParams[string]
	redis token.RdsRepository
}

func New(env *env.EnvParams[string], redis token.RdsRepository) *Service {
	return &Service{
		env:   env,
		redis: redis,
	}
}

func (s *Service) Generate() {}

func (s *Service) Valid() {}
