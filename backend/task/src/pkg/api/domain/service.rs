package domain

import (
	"context"
	"net/url"
	"template/pkg/api/domain/entity"
	"template/pkg/api/domain/repository"
	"template/pkg/api/domain/usecase"
)

type Service struct {
	client repository.Client
}

func NewService(client repository.Client) usecase.Client {
	return &Service{
		client: client,
	}
}

func (s Service) Get(ctx context.Context, path string, query url.Values) ([]*entity.Entity, error) {
	return nil, nil
}
