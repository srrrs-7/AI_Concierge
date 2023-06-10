package token

import (
	"ai_concierge/pkg/entity"
	"ai_concierge/util/env"
	"context"
	"net/url"
)

type Service struct {
	env    *env.EnvParams[string]
	client ApiUseCase
	store  DbUseCase
	redis  RdsUseCase
	sqs    SqsUseCase
	s3     S3UseCase
}

func NewService(
	env *env.EnvParams[string],
	client ApiUseCase,
	store DbUseCase,
	redis RdsUseCase,
	sqs SqsUseCase,
	s3 S3UseCase,
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

func (s Service) Get(ctx context.Context, path string, query url.Values) ([]*entity.Entity, error) {
	return nil, nil
}
