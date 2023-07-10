package auth

import (
	"auth/driver/db/client_master"
	"auth/pkg/domain/master/entity"
	"auth/util/env"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

type UseCase interface {
	Register()
	BasicAuth(req *http.Request) (valid bool, err error)
}

type Service struct {
	env   *env.EnvParams[string]
	store client_master.Store
}

func New(env *env.EnvParams[string], store client_master.Store) *Service {
	return &Service{
		env:   env,
		store: store,
	}
}

func (s *Service) Register() {
	entity := entity.Master{}
}

func (s *Service) BasicAuth(req *http.Request) error {
	header := req.Header.Get("Authorization")
	if header == "" {
		return errors.New("invalid basic auth")
	}
	basic := strings.SplitN(header, " ", 2)
	if len(basic) >= 2 || basic[0] != "Basic" {
		return errors.New("invalid basic auth")
	}
	decoded, err := base64.StdEncoding.DecodeString(basic[1])
	if err != nil {
		return err
	}
	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return errors.New("invalid basic auth")
	}
	id, password := credentials[0], credentials[0]

	master, err := s.store.Find(req.Context(), entity.ClientID(id))
	if err != nil {
		return err
	}
	err = entity.New(master).Compare(entity.Password(password))
	if err != nil {
		return err
	}

	return nil
}
