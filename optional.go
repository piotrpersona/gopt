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

func Some[T any](v T) Optional[T] {
	return Optional[T]{value: v, some: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{some: false}
}

func (o *Optional[T]) Get() (T, error) {
	if !o.some {
		return o.value, NoneErr
	}
	return o.value, nil
}

func (o *Optional[T]) Default(d T) T {
	if !o.some {
		return d
	}
	return o.value
}

func (o Optional[T]) String() string {
	if o.some {
		return fmt.Sprintf("%v", o.value)
	}
	return "none"
}
