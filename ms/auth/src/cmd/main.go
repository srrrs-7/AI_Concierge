package main

import (
	"ai_concierge/driver"
	"ai_concierge/driver/api"
	"ai_concierge/driver/api/domain"
	"ai_concierge/driver/cache"
	"ai_concierge/driver/cache/redis"
	"ai_concierge/driver/db"
	"ai_concierge/driver/db/table"
	"ai_concierge/driver/file"
	awsS3 "ai_concierge/driver/file/s3"
	"ai_concierge/driver/queue"
	awsSqs "ai_concierge/driver/queue/sqs"
	"ai_concierge/pkg"
	"ai_concierge/pkg/domain/auth"
	"ai_concierge/pkg/domain/oauth"
	"ai_concierge/pkg/domain/oidc"
	"ai_concierge/pkg/domain/token"
	"ai_concierge/util/aws"
	"ai_concierge/util/env"
	utilLog "ai_concierge/util/log"
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
