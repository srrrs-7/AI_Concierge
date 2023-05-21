package pkg

import (
	"net/http"
	api "template/pkg/api/domain"
	db "template/pkg/db/table"
	queue "template/pkg/queue/domain"
	sftp "template/pkg/sftp/domain"
)

type Repositories struct {
	apiService   *api.Service
	dbService    *db.Service
	queueService *queue.Service
	sftpService  *sftp.Service
}

func NewRepositories(
	apiService *api.Service,
	dbService *db.Service,
	queueService *queue.Service,
	sftpService *sftp.Service,
) *Repositories {
	return &Repositories{
		apiService:   apiService,
		dbService:    dbService,
		queueService: queueService,
		sftpService:  sftpService,
	}
}

func (repo *Repositories) MainLogic(w http.ResponseWriter, r *http.Request) {

}

func (repo *Repositories) SecondLogic(w http.ResponseWriter, r *http.Request) {

}

func (repo *Repositories) SubLogic(w http.ResponseWriter, r *http.Request) {

}
