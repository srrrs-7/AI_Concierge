package client_master

import (
	"auth/driver/db/client_master/model"
	"auth/pkg/domain/master/entity"
	"context"

	"gorm.io/gorm"
)

type Store interface {
	Find(ctx context.Context, clientID entity.ClientID) (master *model.ClientMaster, err error)
	Create(ctx context.Context, master *model.ClientMaster) error
	Update(ctx context.Context, master *model.ClientMaster, clientID entity.ClientID) error
	Delete(ctx context.Context, clientID entity.ClientID) error
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(ctx context.Context, clientID entity.ClientID) (master *model.ClientMaster, err error) {
	result := r.db.WithContext(ctx).Find(&master, clientID)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (r *Repository) Create(ctx context.Context, master *model.ClientMaster) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&master)
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
func (r *Repository) Update(ctx context.Context, master *model.ClientMaster, clientID entity.ClientID) error {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Model(master).Where("client_id = ?", clientID).Update("is_deleted", 0)
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
func (r *Repository) Delete(ctx context.Context, clientID entity.ClientID) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&model.ClientMaster{}).Where("client_id = ?", clientID).Update("is_deleted", 1)
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
