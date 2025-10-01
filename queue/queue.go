package queue

import "sync"

type Queue[T comparable] interface {
	Enqueue(el T)
	Dequeue() (T, bool)
	Front() (T, bool)
	Clear()
	IsEmpty() bool
	Size() int
}

type myQueue[T comparable] struct {
	s   []T
	mut *sync.Mutex
}

func (q *myQueue[T]) Enqueue(el T) {
	q.mut.Lock()
	defer q.mut.Unlock()

	q.s = append(q.s, el)
}

func (q *myQueue[T]) Dequeue() (T, bool) {
	q.mut.Lock()
	defer q.mut.Unlock()

	if len(q.s) == 0 {
		var zero T
		return zero, false
	}

	el := q.s[0]
	q.s = q.s[1:]

	return el, true
}

func (q *myQueue[T]) Front() (T, bool) {
	q.mut.Lock()
	defer q.mut.Unlock()

	if len(q.s) == 0 {
		var zero T
		return zero, false
	}

	return q.s[0], true
}

func (q *myQueue[T]) Clear() {
	q.mut.Lock()
	defer q.mut.Unlock()

	q.s = []T{}
}

func (q *myQueue[T]) Size() int {
	q.mut.Lock()
	defer q.mut.Unlock()

	return len(q.s)
}

func (q *myQueue[T]) IsEmpty() bool {
	q.mut.Lock()
	defer q.mut.Unlock()

	return len(q.s) == 0
}

func New[T comparable]() Queue[T] {
	return &myQueue[T]{
		s:   []T{},
		mut: &sync.Mutex{},
	}
}
