package env

import (
	"os"
)

type Env struct {
	DB_DRIVER    string
	DB_ADDR      string
	API_URL      string
	HTTP_PORT    string
	AWS_REGION   string
	AWS_S3_BUCKET string
	AWS_SQS_URL   string
}

func SetEnv() *Env {
	return &Env{
		DB_DRIVER:    getEnv("DB_DRIVER", "mysql"),
		DB_ADDR:      getEnv("DB_ADDR", "root:secret@tcp(localhost:3306)/aic?charset=utf8mb4&parseTime=True&loc=Local"),
		API_URL:      getEnv("DOMAIN_URL", "http://localhost:8080"),
		HTTP_PORT:    getEnv("HTTP_PORT", "8888"),
		AWS_REGION:   getEnv("AWS_REGION", "ap-northeast-1"),
		AWS_S3_BUCKET: getEnv("AWS_S3_BUCKET", "local-bucket"),
		AWS_SQS_URL:   getEnv("AWS_SQS_URL", "https://sqs.ap-northeast-1.amazonaws.com/123456789012/my-sqs"),
	}
}

func getEnv(key, value string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return value
}
