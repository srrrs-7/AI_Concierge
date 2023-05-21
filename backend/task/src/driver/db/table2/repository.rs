package table2

import (
	"context"

	"template/driver/db/table2/model"
	"template/pkg/db/domain2/repository"

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

func (r *Repository) Find(ctx context.Context) (model []*model.Model, err error) {

	return model, nil
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
