package main

import (
	"ai_concierge/driver"
	"ai_concierge/driver/api"
	"ai_concierge/driver/api/domain"
	"ai_concierge/driver/cache"
	"ai_concierge/driver/cache/redis"
	"ai_concierge/driver/db"
	"ai_concierge/driver/db/table"
	"ai_concierge/driver/queue"
	awsSqs "ai_concierge/driver/queue/sqs"
	"ai_concierge/driver/sftp"
	awsS3 "ai_concierge/driver/sftp/s3"
	"ai_concierge/pkg"
	apiService "ai_concierge/pkg/api/domain"
	dbService "ai_concierge/pkg/db/table"
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
	queries := table.New(db)
	// aws session
	rdb, err := cache.NewRedis(env)
	if err != nil {
		log.Fatal(err)
	}

	sess := aws.NewAwsSession(env)
	sqs := queue.NewSqs(env, sess)
	s3 := sftp.NewS3(env, sess)
	// repository
	client := domain.NewRepository(env, cli)
	store := table.NewRepository(env, db, queries)
	rdsRepo := redis.NewRepository(env, rdb)
	sqsRepo := awsSqs.NewRepository(env, sqs)
	s3Repo := awsS3.NewRepository(env, s3)
	// service
	apiService := apiService.NewService(env, client)
	dbService := dbService.NewService(env, store)

	// logic
	logicRepo := pkg.NewRepositories(
		apiService,
		dbService,
		rdsRepo,
		sqsRepo,
		s3Repo,
	)

	logger := utilLog.NewRepository()
	// new router
	router := driver.NewRouter(env, logger, logicRepo)
	router.NewRouter()
}
