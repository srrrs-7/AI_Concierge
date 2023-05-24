package domain

import "context"

type Usecase interface {
	Send(ctx context.Context) error
	Receive(ctx context.Context) error
}
