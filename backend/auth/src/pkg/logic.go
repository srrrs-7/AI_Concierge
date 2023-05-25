package pkg

import (
	"fmt"
	"net/http"
	"template/driver/cache/redis"
	sqsRepo "template/driver/queue/sqs"
	s3Repo "template/driver/sftp/s3"
	client "template/pkg/api/domain"
	store "template/pkg/db/table"
)

type Repositories struct {
	client   *client.Service
	store    *store.Service
	rdsRepo      *redis.Repository
	sqsRepo *sqsRepo.Repository
	s3Repo  *s3Repo.Repository
}

func NewRepositories(
	client   *client.Service,
	store    *store.Service,
	rdsRepo      *redis.Repository,
	sqsRepo *sqsRepo.Repository,
	s3Repo  *s3Repo.Repository,
) *Repositories {
	return &Repositories{
		client: client,
		store: store,
		rdsRepo: rdsRepo,
		sqsRepo: sqsRepo,
		s3Repo: s3Repo,
	}
}

func (repo *Repositories) GetToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// oauth api
	// create token
	// set redis
	// insert db

	fmt.Println(ctx)

	w.Write([]byte("ok"))
}

func (repo *Repositories) VerifyToken(w http.ResponseWriter, r *http.Request) {

}
