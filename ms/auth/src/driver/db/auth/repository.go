package auth

import (
	"auth/driver/db/auth/model"
	"auth/pkg/domain/auth/entity"
	"context"

	"gorm.io/gorm"
)

type DbRepository interface {
	Find(ctx context.Context) (authors []*entity.Auth, err error)
	FindByID(ctx context.Context) (author *model.Auth, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context) (auth []*entity.Auth, err error) {
	return nil, nil
}
func (r *Repository) FindByID(ctx context.Context) (auth *model.Auth, err error) {
	result := r.db.Select("user_id", "password", "email", "scope").Where(ctx.Value("userID")).Find(auth)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (r *Repository) Create(ctx context.Context, body map[string]any) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&model.Auth{
			UserID:   "test1",
			Password: "test1",
			Email:    "test1",
			Scope:    "test1",
		})

		if result.Error != nil {
			return result.Error
		}
		return nil
	})

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

func (r *Repository) CreateAuth(ctx context.Context) {

}
