package token

import (
	"auth/pkg/entity"
	"context"
	"net/http"
	"net/url"
)

type ApiRepository interface {
	Get(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Post(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Put(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Delete(ctx context.Context, path string, query url.Values) (*http.Response, error)
}

type DbRepository interface {
	Find(ctx context.Context) (authors []*entity.Auth, err error)
	FindByID(ctx context.Context) (author *entity.Auth, err error)
	Create(ctx context.Context, body map[string]any) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
}

type RdsRepository interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type SqsRepository interface {
	Send(ctx context.Context) error
	Receive(ctx context.Context) error
}

type S3Repository interface {
	Upload(ctx context.Context) error
	Download(ctx context.Context) error
}
