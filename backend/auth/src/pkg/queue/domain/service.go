package domain

import (
	"context"

	"template/util/env"
)

type Service struct {
	env   *env.Env
	queue Usecase
}

func NewService(
	env *env.Env,
	queue Usecase,
) *Service {
	return &Service{
		env:   env,
		queue: queue,
	}
}

func (s *Service) Send(ctx context.Context) error {
	return nil
}
func (s *Service) Receive(ctx context.Context) error {
	return nil
}
