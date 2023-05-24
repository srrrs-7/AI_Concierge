package domain

import (
	"context"
	"net/http"
	"net/url"
	"template/driver/api"

	"template/util/env"
)

type Repository struct {
	env    *env.Env
	client api.Client
}

func NewRepository(
	env *env.Env,
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
