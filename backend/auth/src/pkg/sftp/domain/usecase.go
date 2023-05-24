package domain

import "context"

type UseCase interface {
	Upload(ctx context.Context) error
	Download(ctx context.Context) error
}
