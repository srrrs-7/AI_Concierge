package oauth

import (
	"auth/pkg/entity"
	"auth/util/env"
	"context"
	"net/url"
)

type UseCase interface {
}

type Service struct {
	env    *env.EnvParams[string]
	client ApiRepository
	store  DbRepository
	redis  RdsRepository
	sqs    SqsRepository
	s3     S3Repository
}

func New(
	env *env.EnvParams[string],
	client ApiRepository,
	store DbRepository,
	redis RdsRepository,
	sqs SqsRepository,
	s3 S3Repository,
) *Service {
	return &Service{
		env:    env,
		client: client,
		store:  store,
		redis:  redis,
		sqs:    sqs,
		s3:     s3,
	}
}

func (s Service) Get(ctx context.Context, path string, query url.Values) ([]*entity.Auth, error) {
	return nil, nil
}
