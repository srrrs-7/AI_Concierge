package driver

import (
	"ai_concierge/pkg"
	"ai_concierge/util/env"
	"ai_concierge/util/log"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type Repository struct {
	env       *env.EnvParams[string]
	logger    *log.Repository
	logicRepo *pkg.Repositories
}

func NewRouter(
	env *env.EnvParams[string],
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

	// register
	router.Route("/register", func(c chi.Router) {
		c.Post("/", r.logicRepo.Register)
	})

	// OpenID connect
	router.Route("/oidc", func(c chi.Router) {
		c.Use(r.basicAuth)
		c.Get("/token", r.logicRepo.GetIDToken) // ID tokenの作成
	})

	// OAuth
	router.Route("/oauth/token", func(c chi.Router) {
		c.Get("/create", r.logicRepo.GetAccessToken)    // 認可サーバーでトークン発行
		c.Get("/verify", r.logicRepo.VerifyAccessToken) // トークン認証
	})

	r.logger.Info(fmt.Sprintf("start server on port: %s", r.env.HTTP_PORT.Value))
	err := http.ListenAndServe(":"+r.env.HTTP_PORT.Value, router)
	if err != nil {
		panic(err)
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
