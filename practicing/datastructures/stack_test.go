package datastructures

import (
	"errors"
	"testing"
)

func TestEmptyStack(t *testing.T) {
	t.Run("Empty on Creation", func(t *testing.T) {
		stack := newStack()
		want := true
		got := stack.IsEmpty()
		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})
	t.Run("Not Empty on Addition", func(t *testing.T) {
		stack := newStack()
		stack.Push(1)
		want := false
		got := stack.IsEmpty()
		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})
	t.Run("Empty on Removing Unique Element", func(t *testing.T) {
		stack := newStack()
		stack.Push(1)
		stack.Pop()
		want := true
		got := stack.IsEmpty()
		if got != want {
			t.Errorf("got %t, but want %t", got, want)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Return Correct Value on Pop", func(t *testing.T) {
		stack := newStack()
		stack.Push(1)
		want := 1
		got, err := stack.Pop()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Return Correct Value on Pop", func(t *testing.T) {
		stack := newStack()
		stack.Push(1)
		stack.Push(2)
		stack.Push(4)
		want := 4
		got, err := stack.Pop()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Error on Pop with empty Stack", func(t *testing.T) {
		stack := newStack()
		_, err := stack.Pop()
		if !errors.Is(err, ErrorEmptyStack) {
			t.Errorf("expected %v, but %v", ErrorEmptyStack, err)
		}
	})
}
