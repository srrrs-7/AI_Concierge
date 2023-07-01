package file

import (
	"auth/util/env"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func New(env *env.EnvParams[string], sess *session.Session) *s3.S3 {
	return s3.New(sess)
}
