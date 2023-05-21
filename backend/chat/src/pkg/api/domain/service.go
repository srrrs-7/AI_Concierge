package domain

import (
	"context"
	"net/url"
	"template/pkg/api/domain/entity"
	"template/pkg/api/domain/repository"
	"template/util/env"
)

type Service struct {
	env    *env.Env
	client repository.Client
}

func NewService(
	env *env.Env,
	client repository.Client,
) *Service {
	return &Service{
		env:    env,
		client: client,
	}
}

func (s Service) Get(ctx context.Context, path string, query url.Values) ([]*entity.Entity, error) {
	return nil, nil
}
