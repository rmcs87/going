package datastructures

import "errors"

var (
	ErrorEmptyStack = errors.New("EmptyStack")
)

type Stack struct {
	stack []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrorEmptyStack
	}

	element := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return element, nil
}
