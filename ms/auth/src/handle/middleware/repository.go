package middleware

import (
	"auth/driver/db/client_master"
	"auth/driver/db/client_token"
	"auth/util/env"
	"auth/util/log"
	"net/http"
)

type Middleware interface {
	BasicAuth(next http.Handler) http.Handler
	Certificate(next http.Handler) http.Handler
}

type Repository struct {
	env    *env.EnvParams[string]
	logger *log.Repository
	master client_master.Store
	token  client_token.Store
}

func New(env *env.EnvParams[string], logger *log.Repository, master client_master.Store, token client_token.Store) *Repository {
	return &Repository{
		env:    env,
		logger: logger,
		master: master,
		token:  token,
	}
}

func (r *Repository) BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := r.master.BasicAuth(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		next.ServeHTTP(w, req)
	})
}

func (r *Repository) Certificate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("auth"))
		next.ServeHTTP(w, req)
	})
}
