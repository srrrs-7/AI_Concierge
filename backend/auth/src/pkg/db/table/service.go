package domain

import (
	"context"
	"template/pkg/db/table/entity"
	"template/pkg/db/table/repository"
	"template/util/env"
)

type Service struct {
	env   *env.Env
	store repository.Store
}

func NewService(
	env *env.Env,
	store repository.Store,
) *Service {
	return &Service{
		env:   env,
		store: store,
	}
}

func (s *Service) Find(ctx context.Context) (authors []*entity.Author, err error) {
	return nil, nil
}

func (s *Service) FindByID(ctx context.Context) (author *entity.Author, err error) {
	return nil, nil
}

func (s *Service) Create(ctx context.Context, body map[string]any) error {
	return nil
}

func (s *Service) Update(ctx context.Context) error {
	return nil
}

func (s *Service) Delete(ctx context.Context) error {
	return nil
}
