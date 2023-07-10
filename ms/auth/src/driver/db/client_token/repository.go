package client_token

import (
	"auth/driver/db/client_token/model"
	"auth/pkg/domain/token/entity"
	"context"

	"gorm.io/gorm"
)

type Store interface {
	Find(ctx context.Context, accessToken entity.AccessToken) (token *model.ClientToken, err error)
	Create(ctx context.Context, token *model.ClientToken) error
	Update(ctx context.Context, token *model.ClientToken, accessToken entity.AccessToken) error
	Delete(ctx context.Context, accessToken entity.AccessToken) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context, accessToken entity.AccessToken) (token *model.ClientToken, err error) {
	result := r.db.WithContext(ctx).Find(&token, accessToken)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

func (r *Repository) Create(ctx context.Context, token *model.ClientToken) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&token)
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

func (r *Repository) Update(ctx context.Context, token *model.ClientToken, accessToken entity.AccessToken) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(token).Where("access_token = ?", accessToken).Update("is_deleted", 0)
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

func (r *Repository) Delete(ctx context.Context, accessToken entity.AccessToken) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.ClientToken{}).Where("access_token = ?", accessToken).Update("is_deleted", 1)
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
