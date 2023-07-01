package entity

import "time"

type AccessToken string
type RefreshToken string

type Token struct {
	accessToken  AccessToken
	refreshToken RefreshToken
	expired      time.Time
}

func (t *Token) Expire() bool {
	return t.expired.After(time.Now())
}
