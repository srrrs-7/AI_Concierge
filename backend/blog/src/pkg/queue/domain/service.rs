package domain

import (
	"context"
	"template/pkg/queue/domain/repository"
	"template/pkg/queue/domain/usecase"
)

type Service struct {
	queue repository.Queuer
}

func NewService(queue repository.Queuer) usecase.Queuer {
	return &Service{
		queue: queue,
	}
}

func (s *Service) Send(ctx context.Context) error {
	return nil
}
func (s *Service) Receive(ctx context.Context) error {
	return nil
}
