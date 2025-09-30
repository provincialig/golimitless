package ds

import "sync"

type Stack[T any] interface {
	Push(value T)
	Pop() T
	Peek() T
	Iterate(fn func(value T) bool)
	Clear()
	IsEmpty() bool
	Size() int
}

type myStack[T any] struct {
	s   []T
	mut sync.Mutex
}

func (s *myStack[T]) Push(value T) {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.s = append(s.s, value)
}

func (s *myStack[T]) Pop() T {
	s.mut.Lock()
	defer s.mut.Unlock()

	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val
}

func (s *myStack[T]) Peek() T {
	s.mut.Lock()
	defer s.mut.Unlock()

	return s.s[len(s.s)-1]
}

func (s *myStack[T]) Iterate(fn func(value T) bool) {
	for !s.IsEmpty() {
		if !fn(s.Pop()) {
			return
		}
	}
}

func (s *myStack[T]) Clear() {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.s = []T{}
}

func (s *myStack[T]) Size() int {
	s.mut.Lock()
	defer s.mut.Unlock()

	return len(s.s)
}

func (s *myStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func NewStack[T any]() Stack[T] {
	return &myStack[T]{
		s:   []T{},
		mut: sync.Mutex{},
	}
}
