package awsSqs

import (
	"context"
	"template/pkg/queue/domain/repository"
	"template/util/env"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type Repository struct {
	env    *env.Env
	client *sqs.SQS
}

func NewRepository(
	env *env.Env,
	client *sqs.SQS,
) repository.Queuer {
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
