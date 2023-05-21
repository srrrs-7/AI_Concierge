package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	Fetch(ctx context.Context) (Fetch, error)
}

type Fetch interface {
	Do(ctx context.Context, method Method, path Path, query Query, body Body) (*http.Response, error)
}

type BaseUrl = string
type ID = string
type Secret = string
type Token = string
type Auth struct {
	AuthUrl BaseUrl
	Url     BaseUrl
	ID      ID
	Secret  Secret
	Token   Token
}

type client struct {
	authUrl url.URL
	apiUrl  url.URL
	id      ID
	secret  Secret
}

type Method = string
type Path = string
type Query = url.Values
type Body = io.Reader
type fetch struct {
	token  Token
	apiUrl url.URL
}

func NewClient(auth *Auth) client {
	authUrl, err := url.Parse(string(auth.AuthUrl))
	if err != nil {
		log.Fatal(err)
	}

	apiUrl, err := url.Parse(string(auth.Url))
	if err != nil {
		log.Fatal(err)
	}

	return client{
		authUrl: *authUrl,
		apiUrl:  *apiUrl,
		id:      auth.ID,
		secret:  auth.Secret,
	}
}

// oauth -> basicAuth -> access token -> fetch api
func (c client) Fetch(ctx context.Context) (Fetch, error) {
	authUrl := c.authUrl
	authUrl.Path = "/oauth/token"

	reqBody := url.Values{}
	reqBody.Set("grant_type", "client_credentials")

	req, err := http.NewRequest(
		http.MethodPost,
		authUrl.String(),
		bytes.NewBufferString(reqBody.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(string(c.id), string(c.secret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := newHttpClient().Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		status := resp.StatusCode
		body, _ := io.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprint(status, ":", string(body)))
	}

	body := struct {
		AccessToken Token `json:"access_token"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}

	return &fetch{
		token:  body.AccessToken,
		apiUrl: c.apiUrl,
	}, nil
}

// fetch api execution
func (f fetch) Do(
	ctx context.Context,
	method Method,
	path Path,
	query Query,
	body Body,
) (*http.Response, error) {
	url := f.apiUrl
	url.Path = path
	url.RawQuery = query.Encode()

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", string("Bearer "+f.token))
	req.Header.Set("Content-Type", "application/json")

	return newHttpClient().Do(req.WithContext(ctx))
}

func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}
