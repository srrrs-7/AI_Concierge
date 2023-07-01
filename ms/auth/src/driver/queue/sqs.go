package queue

import (
	"auth/util/env"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func New(env *env.EnvParams[string], sess *session.Session) *sqs.SQS {
	svc := sqs.New(sess)
	return svc
}
