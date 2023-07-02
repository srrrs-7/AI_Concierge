package domain

import (
	"auth/driver/api"
	"auth/util/env"
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

type Repository struct {
	env    *env.EnvParams[string]
	client api.Client
}

func New(
	env *env.EnvParams[string],
	client api.Client,
) *Repository {
	return &Repository{
		env:    env,
		client: client,
	}
}

func (r Repository) Get(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return nil, nil
}

func (r Repository) Post(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return nil, nil
}

func (r Repository) Put(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return nil, nil
}

func (r Repository) Delete(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	return nil, nil
}
