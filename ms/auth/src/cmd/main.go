package main

import (
	"auth/driver/api"
	"auth/driver/api/domain"
	"auth/driver/cache"
	tokenDriver "auth/driver/cache/token"
	"auth/driver/db"
	"auth/driver/file"
	awsS3 "auth/driver/file/s3"
	"auth/driver/queue"
	awsSqs "auth/driver/queue/sqs"
	"auth/handle"
	"auth/handle/middleware"
	"auth/pkg"
	auth "auth/pkg/domain/authorization"
	"auth/pkg/domain/token"
	"auth/util/aws"
	"auth/util/env"
	utilLog "auth/util/log"
	"flag"
	"fmt"
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

	sess := aws.New(env)
	sqsRepo := awsSqs.New(env, queue.New(env, sess))
	s3Repo := awsS3.New(env, file.New(env, sess))

	client := domain.New(env, cli)

	fmt.Println(client, sqsRepo, s3Repo)

	rdsRepo := tokenDriver.New(env, rds)
	authStore := auth.New(gormDb)
	// service
	authService := auth.New(env, authStore)
	tokenService := token.New(env, rdsRepo)

	// middleware
	middleware := middleware.New(env, logger, authService)
	// logic
	logic := pkg.New(
		*authService,
		*tokenService,
	)
	// new router
	router := handle.New(env, logger, middleware, logic)
	router.Start()
}
