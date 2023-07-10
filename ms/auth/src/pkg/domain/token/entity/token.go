package entity

import (
	"auth/driver/db/client_token/model"
	"time"
)

type ClientID string
type AccessToken string
type RefreshToken string
type IsDeleted int

type Token struct {
	ClientID     ClientID     `gorm:"column:client_id" json:"client_id"`
	AccessToken  AccessToken  `gorm:"column:access_token" json:"access_token"`
	RefreshToken RefreshToken `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiresIn    time.Time    `gorm:"column:expires_in" json:"expires_in"`
	CreatedAt    time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time    `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted    IsDeleted    `gorm:"column:is_deleted" json:"is_deleted"`
}

func New(token *model.ClientToken) *Token {
	return &Token{
		ClientID:     ClientID(token.ClientID),
		AccessToken:  AccessToken(token.AccessToken),
		RefreshToken: RefreshToken(token.RefreshToken),
		ExpiresIn:    token.ExpiresIn,
		CreatedAt:    token.CreatedAt,
		UpdatedAt:    token.UpdatedAt,
		IsDeleted:    IsDeleted(token.IsDeleted),
	}
}

func (t *Token) Model() model.ClientToken {
	return model.ClientToken{
		ClientID:     string(t.ClientID),
		AccessToken:  string(t.AccessToken),
		RefreshToken: string(t.RefreshToken),
		ExpiresIn:    t.ExpiresIn,
		CreatedAt:    t.CreatedAt,
		UpdatedAt:    t.UpdatedAt,
		IsDeleted:    int(t.IsDeleted),
	}
}
