package repository

import "context"

type Queuer interface {
	Send(ctx context.Context) error
	Receive(ctx context.Context) error
}
