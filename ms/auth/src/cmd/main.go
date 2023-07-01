package main

import (
	"auth/driver/api"
	"auth/driver/api/domain"
	"auth/driver/cache"
	"auth/driver/cache/redis"
	"auth/driver/db"
	authDriver "auth/driver/db/auth"
	"auth/driver/file"
	awsS3 "auth/driver/file/s3"
	"auth/driver/queue"
	awsSqs "auth/driver/queue/sqs"
	"auth/handle"
	"auth/pkg"
	"auth/pkg/domain/auth"
	"auth/pkg/domain/oauth"
	"auth/pkg/domain/token"
	"auth/util/aws"
	"auth/util/env"
	utilLog "auth/util/log"
	"flag"
)

func init() {
	flag.Parse()
}

func main() {
	logger := utilLog.New()

	env := env.New[string]()
	logger.Info(env)

	// api client
	cli := api.New(&api.Auth{
		AuthUrl: "",
		Url:     "",
		ID:      "",
		Secret:  "",
		Token:   "",
	})
	// DB connection
	gormDb, err := db.New(env)
	if err != nil {
		logger.Error(err)
	}
	defer db.CloseDb(gormDb)

	// redis client
	rds, err := cache.New(env)
	if err != nil {
		logger.Error(err)
	}
	// aws session
	sess := aws.New(env)
	sqs := queue.New(env, sess)
	s3 := file.New(env, sess)
	// repository
	client := domain.New(env, cli)
	rdsRepo := redis.New(env, rds)
	sqsRepo := awsSqs.New(env, sqs)
	s3Repo := awsS3.New(env, s3)
	authRepo := authDriver.New(gormDb)
	// service
	authService := auth.New(env, authRepo, rdsRepo)
	oauthService := oauth.New(env, client, authRepo, rdsRepo, sqsRepo, s3Repo)
	tokenService := token.New(env, client, authRepo, rdsRepo, sqsRepo, s3Repo)
	// logic
	logicRepo := pkg.New(
		*authService,
		*oauthService,
		*tokenService,
	)
	// new router
	router := handle.New(env, logger, logicRepo)
	router.Start()
}
