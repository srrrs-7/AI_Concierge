package env

import (
	"os"
)

type Env struct {
	DbDriver    string
	DbAddr      string
	ApiUrl      string
	HttpPort    string
	AwsRegion   string
	AwsS3Bucket string
	AwsSqsUrl   string
}

func SetEnv() *Env {
	return &Env{
		DbDriver:    getEnv("DB_DRIVER", "mysql"),
		DbAddr:      getEnv("DB_ADDR", "root:secret@tcp(localhost:3306)/aic?charset=utf8mb4&parseTime=True&loc=Local"),
		ApiUrl:      getEnv("DOMAIN_URL", "http://localhost:8080"),
		HttpPort:    getEnv("HTTP_PORT", "8080"),
		AwsRegion:   getEnv("AWS_REGION", "ap-northeast-1"),
		AwsS3Bucket: getEnv("AWS_S3_BUCKET", "local-bucket"),
		AwsSqsUrl:   getEnv("AWS_SQS_URL", "https://sqs.ap-northeast-1.amazonaws.com/123456789012/my-sqs"),
	}
}

func getEnv(key, value string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return value
}
