package domain

import (
	"context"
	"template/pkg/sftp/domain/repository"
	"template/pkg/sftp/domain/usecase"
)

type Service struct {
	sftp repository.SftpHandler
}

func NewService(sftp repository.SftpHandler) usecase.SftpHandler {
	return &Service{
		sftp: sftp,
	}
}

func (s Service) Upload(ctx context.Context) error {
	return nil
}

func (s Service) Download(ctx context.Context) error {
	return nil
}
