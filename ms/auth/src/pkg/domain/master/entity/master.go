package entity

import (
	"auth/driver/db/client_master/model"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ClientID string
type Password string
type Email string
type IsDeleted int

type Master struct {
	ClientID     ClientID  `gorm:"column:client_id" json:"client_id"`
	HashPassword Password  `gorm:"column:hash_password" json:"hash_password"`
	Email        Email     `gorm:"column:email" json:"email"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted    IsDeleted `gorm:"column:is_deleted" json:"is_deleted"`
}

func New(master *model.ClientMaster) *Master {
	return &Master{
		ClientID:     ClientID(master.ClientID),
		HashPassword: Password(master.HashPassword),
		Email:        Email(master.Email),
		CreatedAt:    master.CreatedAt,
		UpdatedAt:    master.UpdatedAt,
		IsDeleted:    IsDeleted(master.IsDeleted),
	}
}

func (m *Master) Model() *model.ClientMaster {
	return &model.ClientMaster{
		ClientID:     string(m.ClientID),
		HashPassword: string(m.HashPassword),
		Email:        string(m.Email),
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		IsDeleted:    int(m.IsDeleted),
	}
}

func (m *Master) EmailValid() bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	isMatch, err := regexp.MatchString(pattern, string(m.Email))
	if err != nil {
		return isMatch
	}
	return isMatch
}

func (m *Master) Encrypt(password Password) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.HashPassword = Password(hash)

	return nil
}

func (m *Master) Compare(password Password) error {
	err := bcrypt.CompareHashAndPassword([]byte(m.HashPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
