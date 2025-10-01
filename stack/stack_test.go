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

	el := s.Pop()
	size := s.Size()

	if el != 3 {
		t.Fatalf("Element: %d", el)
	}

	if size != 2 {
		t.Fatalf("Size: %d", size)
	}

	el = s.Pop()
	size = s.Size()

	if el != 2 {
		t.Fatalf("Element: %d", el)
	}

	if size != 1 {
		t.Fatalf("Size: %d", size)
	}

	el = s.Pop()
	size = s.Size()

	if el != 1 {
		t.Fatalf("Element: %d", el)
	}

	if size != 0 {
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

	el := s.Peek()
	size := s.Size()

	if el != 20 {
		t.Fatalf("Element: %d", el)
	}

	if size != 2 {
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

func Test_Iterate(t *testing.T) {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	sum := 0
	count := 0
	s.Iterate(func(value int) bool {
		sum += value
		count++
		return true
	})

	if sum != 6 {
		t.Fatalf("sum should be 6, got %d", sum)
	}
	if count != 3 {
		t.Fatalf("count should be 3, got %d", count)
	}
	if !s.IsEmpty() {
		t.Fatal("stack should be empty after Iterate")
	}
}
