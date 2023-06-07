package pkg

import (
	"ai_concierge/pkg/domain/auth"
	"ai_concierge/pkg/domain/oauth"
	"ai_concierge/pkg/domain/oidc"
	"ai_concierge/pkg/domain/token"
	"ai_concierge/util/env"
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

}

func (repo *Repository) Oauth(w http.ResponseWriter, r *http.Request) {

}
func (repo *Repository) Oidc(w http.ResponseWriter, r *http.Request) {

}
func (repo *Repository) Token(w http.ResponseWriter, r *http.Request) {

}
