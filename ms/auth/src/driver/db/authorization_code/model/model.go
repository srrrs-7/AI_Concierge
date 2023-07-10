package model

import "time"

type AuthorizationCode struct {
	Code         string    `gorm:"column:code" json:"code"`
	ClientID     string    `gorm:"column:client_id" json:"client_id"`
	ClientSecret string    `gorm:"column:client_secret" json:"client_secret"`
	Scope        string    `gorm:"column:scope" json:"scope"`
	GrantType    string    `gorm:"column:grant_type" json:"grant_type"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted    int       `gorm:"column:is_deleted" json:"is_deleted"`
}

func (a *AuthorizationCode) TableName() string {
	return "authorization_code"
}
