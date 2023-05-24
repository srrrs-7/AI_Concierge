package domain

import (
	"context"

	"template/util/env"
)

type Service struct {
	env  *env.Env
	sftp UseCase
}

func NewService(
	env *env.Env,
	sftp UseCase,
) *Service {
	return &Service{
		env:  env,
		sftp: sftp,
	}
}

func (s Service) Upload(ctx context.Context) error {
	return nil
}

func (s Service) Download(ctx context.Context) error {
	return nil
}
