package model

import "time"

type ClientMaster struct {
	ClientID     string    `gorm:"column:client_id" json:"client_id"`
	HashPassword string    `gorm:"column:hash_password" json:"hash_password"`
	Email        string    `gorm:"column:email" json:"email"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted    int       `gorm:"column:is_deleted" json:"is_deleted"`
}

func (c *ClientMaster) TableName() string {
	return "client_master"
}
