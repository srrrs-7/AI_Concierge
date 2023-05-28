package driver

import (
	"ai_concierge/pkg"
	"ai_concierge/util/env"
	"ai_concierge/util/log"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Repository struct {
	env       *env.Env
	logger    *log.Repository
	logicRepo *pkg.Repositories
}

func NewRouter(
	env *env.Env,
	logger *log.Repository,
	logicRepo *pkg.Repositories,
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

	router.Route("/oauth/token", func(c chi.Router) {
		c.Get("/authorize/", r.logicRepo.GetToken) // 認可サーバーでトークン発行
		c.Get("/verify/", r.logicRepo.VerifyToken) // トークン認証
	})

	r.logger.Info(fmt.Sprintf("start server on port: %s", r.env.HTTP_PORT))
	err := http.ListenAndServe(":"+r.env.HTTP_PORT, router)
	if err != nil {
		panic(err)
	}
}
