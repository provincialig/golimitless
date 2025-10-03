package queue

import (
	"sync"
)

type Queue[T any] interface {
	Enqueue(el T)
	Dequeue() T
	Clear()
	IsEmpty() bool
	Size() int
}

type node[T any] struct {
	value T
	next  *node[T]
}

type linkedListQueue[T any] struct {
	mut  *sync.Mutex
	cond *sync.Cond

	head *node[T]
	tail *node[T]

	size int
}

func (q *linkedListQueue[T]) Enqueue(el T) {
	q.mut.Lock()
	defer q.mut.Unlock()

	newVal := &node[T]{value: el}

	if q.tail == nil {
		q.head = newVal
		q.tail = newVal
	} else {
		q.tail.next = newVal
		q.tail = newVal
	}

	q.size++

	q.cond.Signal()
}

func (q *linkedListQueue[T]) Dequeue() T {
	q.mut.Lock()
	defer q.mut.Unlock()

	for q.size == 0 {
		q.cond.Wait()
	}

	el := q.head

	if el.next != nil {
		q.head = el.next
	} else {
		q.head = nil
		q.tail = nil
	}

	el.next = nil

	q.size--

	return el.value
}

func (q *linkedListQueue[T]) Clear() {
	q.mut.Lock()
	defer q.mut.Unlock()

	q.size = 0
	q.head = nil
	q.tail = nil
}

func (q *linkedListQueue[T]) Size() int {
	q.mut.Lock()
	defer q.mut.Unlock()

	return q.size
}

func (q *linkedListQueue[T]) IsEmpty() bool {
	q.mut.Lock()
	defer q.mut.Unlock()

	return q.size == 0
}

func New[T any]() Queue[T] {
	var mut sync.Mutex
	return &linkedListQueue[T]{
		mut:  &mut,
		cond: sync.NewCond(&mut),
	}
}
