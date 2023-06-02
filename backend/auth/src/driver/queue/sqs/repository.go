package awsSqs

import (
	"ai_concierge/util/env"
	"context"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type Repository struct {
	env    *env.EnvParams[string]
	client *sqs.SQS
}

func NewRepository(
	env *env.EnvParams[string],
	client *sqs.SQS,
) *Repository {
	return &Repository{
		env:    env,
		client: client,
	}
}

func (r *Repository) Send(ctx context.Context) error {
	return nil
}
func (r *Repository) Receive(ctx context.Context) error {
	return nil
}
