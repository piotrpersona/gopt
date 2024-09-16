package gopt

import (
	"errors"
	"fmt"
)

var NoneErr = errors.New("the value is none")

type Optional[T any] struct {
	value T
	some  bool
}

func NewSome[T any](v T) Option[T] {
	return Option[T]{value: v, some: true}
}

func NewNone[T any]() Option[T] {
	return Option[T]{some: false}
}

func (o *Optional[T]) Some() bool {
	return o.some
}

func (o *Optional[T]) Get() (T, error) {
	if !o.some {
		return o.value, NoneErr
	}
	return o.value, nil
}

func (o Optional[T]) String() string {
	if o.some {
		return fmt.Sprintf("%v", o.value)
	}
	return "none"
}
