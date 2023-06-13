package auth

import (
	"auth/driver/db/model"
	"context"

	"gorm.io/gorm"
)

type AuthRepository struct {
	gormDb *gorm.DB
}

func NewRepository(gormDb *gorm.DB) *AuthRepository {
	return &AuthRepository{
		gormDb: gormDb,
	}
}

func (r *AuthRepository) FindByUserID() {}

func (r *AuthRepository) InsertAuth(ctx context.Context) {
	r.gormDb.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
}

func (r *AuthRepository) Validate() {

}
