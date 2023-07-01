package entity

import (
	"auth/driver/db/model"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
)

type UserID string
type Password struct {
	claim string
	hash  string
	key   string
}
type Email string
type Scope string

type Auth struct {
	userID   UserID
	password Password
	email    Email
	scope    Scope
}

func New(auth *model.Auth) *Auth {
	return &Auth{
		userID:   UserID(auth.UserID),
		password: Password{claim: auth.Password, key: auth.Password},
		email:    Email(auth.Email),
		scope:    Scope(auth.Scope),
	}
}

func (a *Auth) ToModel() *model.Auth {
	return &model.Auth{
		UserID:   string(a.userID),
		Password: string(a.password.claim),
		Email:    string(a.email),
		Scope:    string(a.scope),
	}
}

func (a *Auth) ValidEmail() bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	isMatch, err := regexp.MatchString(pattern, string(a.email))
	if err != nil {
		return isMatch
	}
	return isMatch
}

func (a *Auth) Hash() *Auth {
	data := []byte(a.password.claim + a.password.key)
	hash := sha256.Sum256(data)
	hashPass := hex.EncodeToString(hash[:])

	a.password.hash = hashPass
	return a
}
