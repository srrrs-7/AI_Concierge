package repository

import (
	"context"
	"template/driver/db/table"
)

type Store interface {
	Find(ctx context.Context) (authors []*table.Author, err error)
	FindByID(ctx context.Context) (author *table.Author, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}
