package pkg

import (
	"auth/pkg/domain/auth"
	"auth/pkg/domain/oauth"
	"auth/pkg/domain/token"
	"auth/util/env"
	"net/http"
)

type Logic interface {
	Auth(w http.ResponseWriter, req *http.Request)
	Oauth(w http.ResponseWriter, req *http.Request)
	Token(w http.ResponseWriter, req *http.Request)
}

type Repository struct {
	env   *env.EnvParams[string]
	auth  auth.Service
	oauth oauth.Service
	token token.Service
}

func New(auth auth.Service, oauth oauth.Service, token token.Service) *Repository {
	return &Repository{
		auth:  auth,
		oauth: oauth,
		token: token,
	}
}

func (r *Repository) Auth(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("auth"))
}

func (r *Repository) Oauth(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("oauth"))
}
func (r *Repository) Token(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("token"))
}
