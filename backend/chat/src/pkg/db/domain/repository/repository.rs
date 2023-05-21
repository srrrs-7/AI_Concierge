package repository

import (
	"context"
	"template/driver/db/table/model"
)

type Store interface {
	Find(ctx context.Context) (users []model.User, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}
