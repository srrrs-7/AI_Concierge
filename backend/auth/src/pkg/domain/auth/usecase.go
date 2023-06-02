package auth

import (
	"ai_concierge/pkg/entity"
	"context"
	"net/http"
	"net/url"
)

type ApiUseCase interface {
	Get(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Post(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Put(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Delete(ctx context.Context, path string, query url.Values) (*http.Response, error)
}

type DbUseCase interface {
	Find(ctx context.Context) (authors []*entity.Author, err error)
	FindByID(ctx context.Context) (author *entity.Author, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type RdsUseCase interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type SqsUseCase interface {
	Send(ctx context.Context) error
	Receive(ctx context.Context) error
}

type S3UseCase interface {
	Upload(ctx context.Context) error
	Download(ctx context.Context) error
}
