package domain

import (
	"ai_concierge/driver/api"
	"ai_concierge/util/env"
	"context"
	"net/http"
	"net/url"
)

type Repository struct {
	env    *env.EnvParams[string]
	client api.Client
}

func NewRepository(
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
