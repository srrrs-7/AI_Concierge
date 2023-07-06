package entity

import (
	"net/url"
	"time"
)

type Token string
type UserID string
type ClientID string
type Scopes string
type RedirectURL url.URL
type ExpiredAt time.Time

type AuthenticateCode struct {
	token       Token
	userID      UserID
	clientID    ClientID
	scopes      Scopes
	redirectURL RedirectURL
	expiredAt   time.Time
}

const Redirect_URL = "http://localhost:3000"

func New(token Token, userID UserID, clientID ClientID, scopes Scopes, redirectURL RedirectURL, expiredAt time.Time) *AuthenticateCode {
	return &AuthenticateCode{
		token:       token,
		userID:      userID,
		clientID:    clientID,
		scopes:      scopes,
		redirectURL: redirectURL,
		expiredAt:   expiredAt,
	}
}
