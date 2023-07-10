package authorization_code

import (
	"auth/driver/db/authorization_code/model"
	"context"

	"gorm.io/gorm"
)

type Store interface {
	Find(ctx context.Context, authCode entity.Code) (code *model.AuthorizationCode, err error)
	Create(ctx context.Context, code *model.AuthorizationCode) error
	Update(ctx context.Context, code *model.AuthorizationCode, authCode entity.Code) error
	Delete(ctx context.Context, authCode entity.Code) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context, authCode entity.Code) (code *model.AuthorizationCode, err error) {
	result := r.db.WithContext(ctx).Find(&code, authCode)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (r *Repository) Create(ctx context.Context, code *model.AuthorizationCode) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&code)
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
func (r *Repository) Update(ctx context.Context, code *model.AuthorizationCode, authCode entity.Code) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(code).Where("code = ?", authCode).Update("is_deleted", 0)
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
func (r *Repository) Delete(ctx context.Context, authCode entity.Code) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.AuthorizationCode{}).Where("code = ?", authCode).Update("is_deleted", 1)
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
