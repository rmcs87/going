package datastructures

import (
	"errors"
)

var (
	ErrorEmptyStack = errors.New("EmptyStack")
)

type Stack[T any] struct {
	stack []T
}

func newStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrorEmptyStack
	}

	element := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return element, nil
}
