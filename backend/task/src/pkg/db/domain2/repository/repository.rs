package repository

import (
	"context"
	"template/driver/db/table2/model"
)

type Store interface {
	Find(ctx context.Context) (model []*model.Model, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}
