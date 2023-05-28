package pkg

import (
	"ai_concierge/driver/cache/redis"
	sqsRepo "ai_concierge/driver/queue/sqs"
	s3Repo "ai_concierge/driver/sftp/s3"
	client "ai_concierge/pkg/api/domain"
	store "ai_concierge/pkg/db/table"
	"fmt"
	"net/http"
)

type Repositories struct {
	client  *client.Service
	store   *store.Service
	rdsRepo *redis.Repository
	sqsRepo *sqsRepo.Repository
	s3Repo  *s3Repo.Repository
}

func NewRepositories(
	client *client.Service,
	store *store.Service,
	rdsRepo *redis.Repository,
	sqsRepo *sqsRepo.Repository,
	s3Repo *s3Repo.Repository,
) *Repositories {
	return &Repositories{
		client:  client,
		store:   store,
		rdsRepo: rdsRepo,
		sqsRepo: sqsRepo,
		s3Repo:  s3Repo,
	}
}

func (repo *Repositories) Register(w http.ResponseWriter, r *http.Request) {
	// ID, password register
}

func (repo *Repositories) GetIDToken(w http.ResponseWriter, r *http.Request) {

	// oauth api
	// create token
	// set redis
	// insert db

	w.Write([]byte("ok"))
}

func (repo *Repositories) GetAccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// oauth api
	// create token
	// set redis
	// insert db

	fmt.Println(ctx)

	w.Write([]byte("ok"))
}

func (repo *Repositories) VerifyAccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println(ctx)
}
