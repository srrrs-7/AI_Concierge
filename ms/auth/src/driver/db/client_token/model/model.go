package model

import "time"

type ClientToken struct {
	ClientID     string    `gorm:"column:client_id" json:"client_id"`
	AccessToken  string    `gorm:"column:access_token" json:"access_token"`
	RefreshToken string    `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiresIn    time.Time `gorm:"column:expires_in" json:"expires_in"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted    int       `gorm:"column:is_deleted" json:"is_deleted"`
}

func (c *ClientToken) TableName() string {
	return "client_token"
}
