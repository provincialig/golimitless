package stack_test

import (
	"testing"

	"github.com/provincialig/golimitless/stack"
)

func Test_Push(t *testing.T) {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)

	size := s.Size()
	if size != 2 {
		t.Fatalf("Size: %d", size)
	}

	if s.IsEmpty() {
		t.Fatal("Stack must be enmpty")
	}
}

func Test_Pop(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if el, ok := s.Pop(); !ok || el != 3 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 2 {
		t.Fatalf("Size: %d", size)
	}

	if el, ok := s.Pop(); !ok || el != 2 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 1 {
		t.Fatalf("Size: %d", size)
	}

	if el, ok := s.Pop(); !ok || el != 1 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 0 {
		t.Fatalf("Size: %d", size)
	}

	if !s.IsEmpty() {
		t.Fatal("Is empty", s.IsEmpty())
	}
}

func Test_Peek(t *testing.T) {
	s := stack.New[int]()
	s.Push(10)
	s.Push(20)

	if el, ok := s.Peek(); !ok || el != 20 {
		t.Fatalf("Element: %d", el)
	}

	if size := s.Size(); size != 2 {
		t.Fatalf("Size: %d", size)
	}
}

func Test_Clear(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	initialSize := s.Size()
	if initialSize != 3 {
		t.Fatalf("Initial size: %d", initialSize)
	}

	s.Clear()

	finalSize := s.Size()
	if finalSize != 0 {
		t.Fatalf("Final size: %d", initialSize)
	}
}
