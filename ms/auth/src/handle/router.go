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
	logicRepo  *pkg.Repository
}

func New(env *env.EnvParams[string], logger *log.Repository, middleware *middleware.Repository, logicRepo *pkg.Repository) *Repository {
	return &Repository{
		env:        env,
		logger:     logger,
		middleware: middleware,
		logicRepo:  logicRepo,
	}
}

func (r *Repository) Start() {
	router := chi.NewRouter()

	router.Use(log.NewLogging)
	router.Use(r.middleware.BasicAuth)
	router.Use(r.middleware.Authenticator)

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
