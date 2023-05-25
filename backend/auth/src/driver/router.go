package driver

import (
	"fmt"
	"net/http"
	"template/pkg"
	"template/util/env"
	"template/util/log"

	"github.com/go-chi/chi"
)

type Repository struct {
	env       *env.Env
	logicRepo *pkg.Repositories
}

func NewRouter(
	env *env.Env,
	logicRepo *pkg.Repositories,
) *Repository {
	return &Repository{
		env:       env,
		logicRepo: logicRepo,
	}
}

func (r *Repository) NewRouter() {
	router := chi.NewRouter()
	router.Use(log.NewLogging)

	router.Route("/oauth/token", func(c chi.Router) {
		c.Get("/get/", r.logicRepo.GetToken)
		c.Get("/verify/", r.logicRepo.VerifyToken)
	})

	fmt.Println("start server on port:", r.env.HTTP_PORT)
	err := http.ListenAndServe(":"+r.env.HTTP_PORT, router)
	if err != nil {
		panic(err)
	}
}
