package stack

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrTimeout  = errors.New("context timeout")
	ErrCanceled = errors.New("context canceled")
)

type Stack[T any] interface {
	Push(value T)
	TryPop() (T, bool)
	Pop(ctx context.Context) (T, error)
	TryPeek() (T, bool)
	Peek(ctx context.Context) (T, error)
	Clear()
	IsEmpty() bool
	Size() int
}

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedListStack[T any] struct {
	mut  *sync.Mutex
	cond *sync.Cond

	top *node[T]

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

func (s *linkedListStack[T]) popUnsafe() (T, bool) {
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

func (s *linkedListStack[T]) TryPop() (T, bool) {
	s.mut.Lock()
	defer s.mut.Unlock()

	return s.popUnsafe()
}

func (s *linkedListStack[T]) Pop(ctx context.Context) (T, error) {
	if ctx.Err() != nil {
		var zero T
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return zero, ErrTimeout
		}
		return zero, ErrCanceled
	}

	s.mut.Lock()
	defer s.mut.Unlock()

	ctxDone := make(chan struct{})

	go func() {
		<-ctx.Done()

		close(ctxDone)

		s.mut.Lock()
		s.cond.Broadcast()
		s.mut.Unlock()
	}()

	for {
		if item, ok := s.popUnsafe(); ok {
			return item, nil
		}

		select {
		case <-ctxDone:
			var zero T
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return zero, ErrTimeout
			}
			return zero, ErrCanceled
		default:
		}

		s.cond.Wait()
	}
}

func (s *linkedListStack[T]) peekUnsafe() (T, bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}

	return s.top.value, true
}

func (s *linkedListStack[T]) TryPeek() (T, bool) {
	s.mut.Lock()
	defer s.mut.Unlock()

	return s.peekUnsafe()
}

func (s *linkedListStack[T]) Peek(ctx context.Context) (T, error) {
	if ctx.Err() != nil {
		var zero T
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return zero, ErrTimeout
		}
		return zero, ErrCanceled
	}

	s.mut.Lock()
	defer s.mut.Unlock()

	ctxDone := make(chan struct{})

	go func() {
		<-ctx.Done()

		close(ctxDone)

		s.mut.Lock()
		s.cond.Broadcast()
		s.mut.Unlock()
	}()

	for {
		if item, ok := s.peekUnsafe(); ok {
			return item, nil
		}

		select {
		case <-ctxDone:
			var zero T
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return zero, ErrTimeout
			}
			return zero, ErrCanceled
		default:
		}

		s.cond.Wait()
	}
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
	var mut sync.Mutex
	return &linkedListStack[T]{
		mut:  &mut,
		cond: sync.NewCond(&mut),
	}
}
