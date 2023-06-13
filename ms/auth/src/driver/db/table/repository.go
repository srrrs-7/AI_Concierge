package table

import (
	"auth/pkg/entity"
	"auth/util/env"
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	env *env.EnvParams[string]
	db  *gorm.DB
}

func NewRepository(
	env *env.EnvParams[string],
	db *gorm.DB,
) *Repository {
	return &Repository{
		env: env,
		db:  db,
	}
}

func (r *Repository) Find(ctx context.Context) (authors []*entity.Author, err error) {
	return nil, nil
}

func (r *Repository) FindByID(ctx context.Context) (author *entity.Author, err error) {
	return nil, nil
}

func (r *Repository) Create(ctx context.Context, body map[string]any) error {
	return nil
}

func (r *Repository) Update(ctx context.Context) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context) error {
	return nil
}
