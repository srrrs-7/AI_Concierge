package env

import (
	"fmt"
	"os"
)

type Value interface {
	~string
}

type EnvParams[T Value] struct {
	DB_DRIVER     Env[T]
	DB_ADDR       Env[T]
	API_URL       Env[T]
	API_PORT      Env[T]
	REDIS_URL     Env[T]
	AWS_REGION    Env[T]
	AWS_S3_BUCKET Env[T]
	AWS_SQS_URL   Env[T]
}

type Env[T Value] struct {
	Name  string
	Value T
}

func (e Env[T]) Output() string {
	return fmt.Sprintf("%s=%s", e.Name, e.Value)
}

func newEnv[T Value](name string) Env[T] {
	return Env[T]{
		Name:  name,
		Value: T(os.Getenv(name)),
	}
}

func SetEnv[T Value]() *EnvParams[T] {
	return &EnvParams[T]{
		DB_DRIVER:     newEnv[T]("DB_DRIVER"),
		DB_ADDR:       newEnv[T]("DB_ADDR"),
		API_URL:       newEnv[T]("API_URL"),
		API_PORT:      newEnv[T]("API_PORT"),
		REDIS_URL:     newEnv[T]("REDIS_URL"),
		AWS_REGION:    newEnv[T]("AWS_REGION"),
		AWS_S3_BUCKET: newEnv[T]("AWS_S3_BUCKET"),
		AWS_SQS_URL:   newEnv[T]("AWS_SQS_URL"),
	}
}
