package queue

import (
	"ai_concierge/util/env"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewSqs(env *env.Env, sess *session.Session) *sqs.SQS {
	svc := sqs.New(sess)
	return svc
}
