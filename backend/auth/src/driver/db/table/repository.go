package table

import (
	"ai_concierge/pkg/entity"
	"ai_concierge/util/env"
	"context"
	"database/sql"
)

type Repository struct {
	env     *env.EnvParams[string]
	db      *sql.DB
	queries *Queries
}

func NewRepository(
	env *env.EnvParams[string],
	db *sql.DB,
	queries *Queries,
) *Repository {
	return &Repository{
		env:     env,
		db:      db,
		queries: queries,
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
