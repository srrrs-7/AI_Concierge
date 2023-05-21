package domain

import (
	"context"
	"template/pkg/queue/domain/repository"
	"template/util/env"
)

type Service struct {
	env   *env.Env
	queue repository.Queuer
}

func NewService(
	env *env.Env,
	queue repository.Queuer,
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
