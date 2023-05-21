package aws

import (
	"log"
	"template/util/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAwsSession(env *env.Env) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AwsRegion),
	})
	if err != nil {
		log.Fatal(err)
	}
	return sess
}
