package entity

import "time"

type AccessToken string
type RefreshToken string
type UserID string
type ClientID string

type Token struct {
	accessToken  AccessToken
	refreshToken RefreshToken
	userID       UserID
	clientID     ClientID
	expired      time.Time
}

func (t *Token) Expire() bool {
	return t.expired.After(time.Now())
}
