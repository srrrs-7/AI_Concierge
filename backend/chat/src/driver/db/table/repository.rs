package table

import (
	"context"

	"template/driver/db/table/model"
	"template/pkg/db/domain/repository"

	"template/util/env"

	"gorm.io/gorm"
)

type Repository struct {
	env *env.Env
	tx  *gorm.DB
}

func NewRepository(
	env *env.Env,
	tx *gorm.DB,
) repository.Store {
	return &Repository{
		env: env,
		tx:  tx,
	}
}

func (r *Repository) Find(ctx context.Context) (users []model.User, err error) {

	return users, nil
}

func (r *Repository) Create(ctx context.Context, body map[string]any) error {

	err := r.tx.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			d := tx.Create(body)
			if d.Error != nil {
				return d.Error
			}
			return nil
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context) error {
	return nil
}
