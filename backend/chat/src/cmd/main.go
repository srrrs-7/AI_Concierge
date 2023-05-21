package main

import (
	"flag"
	"log"
	"template/driver"
	"template/driver/api"
	"template/driver/api/domain"
	"template/driver/db"
	"template/driver/db/table"
	"template/driver/queue"
	awsSqs "template/driver/queue/sqs"
	"template/driver/sftp"
	awsS3 "template/driver/sftp/s3"
	"template/pkg"
	apiService "template/pkg/api/domain"
	dbService "template/pkg/db/table"
	queueService "template/pkg/queue/domain"
	sftpService "template/pkg/sftp/domain"
	"template/util/aws"
	"template/util/env"
)

func init() {
	flag.Parse()
}

func main() {
	env := env.SetEnv()
	// api client
	cli := api.NewClient(&api.Auth{
		AuthUrl: "",
		Url:     "",
		ID:      "",
		Secret:  "",
		Token:   "",
	})
	// DB connection
	db, err := db.NewDb(env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// sqlc queries
	queries := table.New(db)
	// aws session
	sess := aws.NewAwsSession(env)
	sqs := queue.NewSqs(env, sess)
	s3 := sftp.NewS3(env, sess)
	// repository
	client := domain.NewRepository(env, cli)
	store := table.NewRepository(env, db, queries)
	queue := awsSqs.NewRepository(env, sqs)
	sftp := awsS3.NewRepository(env, s3)
	// service
	apiService := apiService.NewService(env, client)
	dbService := dbService.NewService(env, store)
	queueService := queueService.NewService(env, queue)
	sftpService := sftpService.NewService(env, sftp)

	// logic
	logicRepo := pkg.NewRepositories(
		apiService,
		dbService,
		queueService,
		sftpService,
	)

	// new router
	driver.NewRouter(env, *logicRepo)
}
