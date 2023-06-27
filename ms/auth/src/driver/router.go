package driver

import (
	"auth/pkg"
	"auth/util/env"
	"auth/util/log"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type Repository struct {
	env       *env.EnvParams[string]
	logger    *log.Repository
	logicRepo *pkg.Repository
}

func NewRouter(
	env *env.EnvParams[string],
	logger *log.Repository,
	logicRepo *pkg.Repository,
) *Repository {
	return &Repository{
		env:       env,
		logger:    logger,
		logicRepo: logicRepo,
	}
}

func (r *Repository) NewRouter() {
	router := chi.NewRouter()
	router.Use(log.NewLogging)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/oauth", func(c chi.Router) {
		c.Get("/authorize", r.logicRepo.Auth)
		c.Get("/callback", r.logicRepo.Auth)
		c.Post("/token", r.logicRepo.Auth)
		c.Get("/authenticate", r.logicRepo.Auth)
	})

	r.logger.Info(fmt.Sprintf("start server on port: %s", r.env.API_PORT.Value))
	err := http.ListenAndServe(":"+r.env.API_PORT.Value, router)
	if err != nil {
		r.logger.Error(err)
	}
}

func (r *Repository) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		auth := strings.SplitN(authHeader, " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		credentials, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			http.Error(w, "Failed to decode authorization header", http.StatusUnauthorized)
			return
		}

		// credentialsには"username:password"の形式で認証情報が入っています
		// ここでcredentialsを適切に検証する処理を実装してください

		fmt.Println(string(credentials))

		next.ServeHTTP(w, r)
	})
}
