package domain2

import (
	"context"
	"template/pkg/db/domain/entity"
	"template/pkg/db/domain2/repository"
	"template/pkg/db/domain2/usecase"
)

type Service struct {
	store repository.Store
}

func NewService(
	store repository.Store,
) usecase.Store {
	return &Service{
		store: store,
	}
}

func (s Service) Find(ctx context.Context) ([]*entity.Entity, error) {
	return nil, nil
}
