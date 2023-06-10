package pkg

import (
	"auth/pkg/domain/auth"
	"auth/pkg/domain/oauth"
	"auth/pkg/domain/oidc"
	"auth/pkg/domain/token"
	"auth/util/env"
	"net/http"
)

type Repository struct {
	env   *env.EnvParams[string]
	auth  auth.Service
	oidc  oidc.Service
	oauth oauth.Service
	token token.Service
}

func NewRepository(
	env *env.EnvParams[string],
	auth auth.Service,
	oidc oidc.Service,
	oauth oauth.Service,
	token token.Service,
) *Repository {
	return &Repository{
		env:   env,
		auth:  auth,
		oidc:  oidc,
		oauth: oauth,
		token: token,
	}
}

func (repo *Repository) Auth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth"))
}

func (repo *Repository) Oauth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oauth"))
}
func (repo *Repository) Oidc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oidc"))
}
func (repo *Repository) Token(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("token"))
}
