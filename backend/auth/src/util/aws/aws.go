package aws

import (
	"ai_concierge/util/env"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAwsSession(env *env.Env) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AWS_REGION),
	})
	if err != nil {
		log.Fatal(err)
	}
	return sess
}
