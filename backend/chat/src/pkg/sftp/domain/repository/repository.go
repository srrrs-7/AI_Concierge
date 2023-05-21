package repository

import "context"

type SftpHandler interface {
	Upload(ctx context.Context) error
	Download(ctx context.Context) error
}
