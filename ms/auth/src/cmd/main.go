package main

import (
	"auth/driver"
	"auth/driver/api"
	"auth/driver/api/domain"
	"auth/driver/cache"
	"auth/driver/cache/redis"
	"auth/driver/db"
	"auth/driver/db/table"
	"auth/driver/file"
	awsS3 "auth/driver/file/s3"
	"auth/driver/queue"
	awsSqs "auth/driver/queue/sqs"
	"auth/pkg"
	"auth/pkg/domain/auth"
	"auth/pkg/domain/oauth"
	"auth/pkg/domain/oidc"
	"auth/pkg/domain/token"
	"auth/util/aws"
	"auth/util/env"
	utilLog "auth/util/log"
	"flag"
	"log"
)

func init() {
	flag.Parse()
}

func main() {
	env := env.SetEnv[string]()
	log.Println(env)

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
	queries := table.New(db)

	// redis client
	rds, err := cache.NewRedis(env)
	if err != nil {
		log.Fatal(err)
	}
	// aws session
	sess := aws.NewAwsSession(env)
	sqs := queue.NewSqs(env, sess)
	s3 := file.NewS3(env, sess)
	// repository
	client := domain.NewRepository(env, cli)
	store := table.NewRepository(env, db, queries)
	rdsRepo := redis.NewRepository(env, rds)
	sqsRepo := awsSqs.NewRepository(env, sqs)
	s3Repo := awsS3.NewRepository(env, s3)
	// service
	authService := auth.NewService(env, store, rdsRepo)
	oauthService := oauth.NewService(env, client, store, rdsRepo, sqsRepo, s3Repo)
	oidcService := oidc.NewService(env, client, store, rdsRepo, sqsRepo, s3Repo)
	tokenService := token.NewService(env, client, store, rdsRepo, sqsRepo, s3Repo)

	// logic
	logicRepo := pkg.NewRepository(
		env,
		*authService,
		*oidcService,
		*oauthService,
		*tokenService,
	)

	logger := utilLog.NewRepository()
	// new router
	router := driver.NewRouter(env, logger, logicRepo)
	router.NewRouter()
}
