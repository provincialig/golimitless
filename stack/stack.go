package stack

import "sync"

type Stack[T any] interface {
	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
	Clear()
	IsEmpty() bool
	Size() int
}

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedListStack[T any] struct {
	mut  sync.Mutex
	top  *node[T]
	size int
}

func (s *linkedListStack[T]) Push(value T) {
	s.mut.Lock()
	defer s.mut.Unlock()

	newVal := &node[T]{value: value}

	if s.top == nil {
		s.top = newVal
	} else {
		newVal.next = s.top
		s.top = newVal
	}

	s.size++
}

func (s *linkedListStack[T]) Pop() (T, bool) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if s.top == nil {
		var zero T
		return zero, false
	}

	el := s.top
	s.top = el.next
	el.next = nil

	s.size--

	return el.value, true
}

func (s *linkedListStack[T]) Peek() (T, bool) {
	s.mut.Lock()
	defer s.mut.Unlock()

	if s.top == nil {
		var zero T
		return zero, false
	}

	return s.top.value, true
}

func (s *linkedListStack[T]) Clear() {
	s.mut.Lock()
	defer s.mut.Unlock()

	s.top = nil
	s.size = 0
}

func (s *linkedListStack[T]) Size() int {
	s.mut.Lock()
	defer s.mut.Unlock()

	return s.size
}

func (s *linkedListStack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func New[T any]() Stack[T] {
	return &linkedListStack[T]{}
}
