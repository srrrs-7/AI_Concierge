package middleware

import (
	"auth/pkg/domain/auth"
	"auth/util/env"
	"auth/util/log"
	"net/http"
)

type Middleware interface {
	BasicAuth(next http.Handler) http.Handler
	Authenticator(next http.Handler) http.Handler
}

type Repository struct {
	env    *env.EnvParams[string]
	logger *log.Repository
	auth   auth.Service
}

func New(env *env.EnvParams[string], logger *log.Repository, auth auth.Service) *Repository {
	return &Repository{
		env:    env,
		logger: logger,
		auth:   auth,
	}
}

func (r *Repository) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := r.auth.Auth(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		next.ServeHTTP(w, req)
	})
}

func (r *Repository) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("auth"))
		next.ServeHTTP(w, req)
	})
}
