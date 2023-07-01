package auth

import (
	"auth/pkg/domain/auth/entity"
	"auth/util/env"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

type UseCase interface {
}

type Service struct {
	env   *env.EnvParams[string]
	store DbRepository
	redis RdsRepository
}

func New(env *env.EnvParams[string], store DbRepository, redis RdsRepository) *Service {
	return &Service{
		env:   env,
		store: store,
		redis: redis,
	}
}

func (s *Service) Auth(req *http.Request) error {
	ctx := req.Context()
	auth, err := s.store.FindByID(ctx)
	if err != nil {
		return err
	}
	entity := entity.New(auth).Hash()
	if entity.ValidEmail() {
		return nil
	}

	id, password := basicAuth(req)
	if auth.UserID != id {
		return fmt.Errorf("basic auth id invalid: %s", id)
	} else if auth.Password != password {
		return fmt.Errorf("basic auth password invalid: %s", password)
	}

	return nil
}

func basicAuth(req *http.Request) (id string, password string) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return id, password
	}

	auth := strings.SplitN(authHeader, " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return id, password
	}

	decoded, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return id, password
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return id, password
	}
	id, password = credentials[0], credentials[0]

	return
}

func (s *Service) Valid() {

}
