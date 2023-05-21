package repository

import (
	"context"
	"net/http"
	"net/url"
)

type Client interface {
	Get(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Post(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Put(ctx context.Context, path string, query url.Values) (*http.Response, error)
	Delete(ctx context.Context, path string, query url.Values) (*http.Response, error)
}
