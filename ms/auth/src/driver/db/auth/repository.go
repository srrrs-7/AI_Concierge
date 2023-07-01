package auth

import (
	"auth/driver/db/model"
	"auth/pkg/entity"
	"context"

	"gorm.io/gorm"
)

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
