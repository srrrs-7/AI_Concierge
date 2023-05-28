package domain

import (
	"ai_concierge/pkg/db/table/entity"
	"context"
)

type UseCase interface {
	Find(ctx context.Context) (authors []*entity.Author, err error)
	FindByID(ctx context.Context) (author *entity.Author, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}
