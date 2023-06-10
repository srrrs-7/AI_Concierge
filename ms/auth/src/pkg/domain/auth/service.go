package auth

import (
	"auth/util/env"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
)

type Service struct {
	env   *env.EnvParams[string]
	store DbUseCase
	redis RdsUseCase
}

func NewService(
	env *env.EnvParams[string],
	store DbUseCase,
	redis RdsUseCase,
) *Service {
	return &Service{
		env:   env,
		store: store,
		redis: redis,
	}
}

func (s *Service) BasicAuth(r *http.Request) (id, password string, err error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", "", errors.New("missing authorization header")
	}

	auth := strings.SplitN(authHeader, " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return "", "", errors.New("invalid authorization header")
	}

	decoded, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return "", "", errors.New("invalid authorization header")
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return "", "", errors.New("invalid authorization header")
	}

	return credentials[0], credentials[1], nil
}

func (s *Service) Hash(id, password string) (string, string) {

	hash := sha256.Sum256([]byte(password))
	hashString := hex.EncodeToString(hash[:])

	return id, hashString

}
