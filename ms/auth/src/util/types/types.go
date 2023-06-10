package types

import "errors"

type Option[T any] struct {
	Value T
	Valid bool
}

func NewOption[T any]() *Option[T] {
	return &Option[T]{}
}

func (o *Option[T]) Some(value T) *Option[T] {
	return &Option[T]{
		Value: value,
		Valid: true,
	}
}

func (o *Option[T]) None(value T) *Option[T] {
	return &Option[T]{
		Value: value,
		Valid: false,
	}
}

type Result[T any] struct {
	Valid bool
	Error error
}

func NewResult[T any]() *Result[T] {
	return &Result[T]{}
}

func (r *Result[T]) Ok() *Result[T] {
	return &Result[T]{
		Valid: true,
		Error: nil,
	}
}

func (r *Result[T]) Err() *Result[T] {
	return &Result[T]{
		Valid: false,
		Error: errors.New("error"),
	}
}
