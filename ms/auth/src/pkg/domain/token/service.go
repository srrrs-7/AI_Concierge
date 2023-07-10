package token

import (
	"auth/driver/cache/token"
	"auth/pkg/domain/token/entity"
	"auth/util/env"
	"crypto/rand"
	"encoding/base64"
	"time"
)

type UseCase interface {
}

type Service struct {
	env   *env.EnvParams[string]
	redis token.RdsRepository
}

const TOKEN_LENGTH = 64

func New(env *env.EnvParams[string], redis token.RdsRepository) *Service {
	return &Service{
		env:   env,
		redis: redis,
	}
}

func (s *Service) AccessToken() (accessToken entity.AccessToken, err error) {
	b := make([]byte, TOKEN_LENGTH)
	_, err = rand.Read(b)
	accessToken = entity.AccessToken(base64.URLEncoding.EncodeToString(b)[:TOKEN_LENGTH])
	if err != nil {
		return "", err
	}
	return
}

func (s *Service) RefreshToken() (refreshToken entity.RefreshToken, err error) {
	b := make([]byte, TOKEN_LENGTH)
	_, err = rand.Read(b)
	refreshToken = entity.RefreshToken(base64.URLEncoding.EncodeToString(b)[:TOKEN_LENGTH])
	if err != nil {
		return "", err
	}
	return
}

func (s *Service) Expire() (expiresIn time.Time) {
	expiresIn = time.Now().Add(time.Hour)
	return
}
