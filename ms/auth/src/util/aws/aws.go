package aws

import (
	"auth/util/env"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAwsSession(env *env.EnvParams[string]) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AWS_REGION.Value),
	})
	if err != nil {
		log.Fatal(err)
	}
	return sess
}
