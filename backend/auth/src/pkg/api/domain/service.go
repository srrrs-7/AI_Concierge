package domain

import (
	"ai_concierge/pkg/api/domain/entity"
	"ai_concierge/util/env"
	"context"
	"net/url"
)

type Service struct {
	env    *env.Env
	client UseCase
}

func NewService(
	env *env.Env,
	client UseCase,
) *Service {
	return &Service{
		env:    env,
		client: client,
	}
}

func (s Service) Get(ctx context.Context, path string, query url.Values) ([]*entity.Entity, error) {
	return nil, nil
}
