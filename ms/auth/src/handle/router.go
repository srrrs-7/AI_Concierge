package handle

import (
	"auth/handle/middleware"
	"auth/pkg"
	"auth/util/env"
	"auth/util/log"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Repository struct {
	env        *env.EnvParams[string]
	logger     *log.Repository
	middleware middleware.Middleware
	logic      *pkg.Repository
}

func New(env *env.EnvParams[string], logger *log.Repository, middleware *middleware.Repository, logic *pkg.Repository) *Repository {
	return &Repository{
		env:        env,
		logger:     logger,
		middleware: middleware,
		logic:      logic,
	}
}

func (r *Repository) Start() {
	router := chi.NewRouter()

	router.Use(log.NewLogging)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// master register
	router.Route("/master", func(c chi.Router) {
		c.Post("/register", r.logic.Auth)
		c.Post("/login", r.logic.Auth)
	})

	// issue auth token
	router.Route("/auth", func(c chi.Router) {
		c.Use(r.middleware.Authenticate)
		c.Post("/authenticate", r.logic.Auth)
	})

	// issue certificate token
	router.Route("/oauth", func(c chi.Router) {
		c.Use(r.middleware.BasicAuth)
		c.Use(r.middleware.Certificate)

		c.Post("/certificate", r.logic.Auth)
	})

	r.logger.Info(fmt.Sprintf("start server on port: %s", r.env.API_PORT.Value))
	err := http.ListenAndServe(":"+r.env.API_PORT.Value, router)
	if err != nil {
		r.logger.Error(err)
	}
}
