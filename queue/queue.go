package queue

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrTimeout  = errors.New("context timeout")
	ErrCanceled = errors.New("context canceled")
)

type Queue[T any] interface {
	Enqueue(el T)
	TryDequeue() (T, bool)
	Dequeue(ctx context.Context) (T, error)
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

	size int64
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

	atomic.AddInt64(&q.size, 1)

	q.cond.Signal()
}

func (q *linkedListQueue[T]) dequeueUnsafe() (T, bool) {
	if q.size == 0 {
		var zero T
		return zero, false
	}

	el := q.head

	if el.next != nil {
		q.head = el.next
	} else {
		q.head = nil
		q.tail = nil
	}

	el.next = nil

	atomic.AddInt64(&q.size, -1)

	return el.value, true
}

func (q *linkedListQueue[T]) TryDequeue() (T, bool) {
	q.mut.Lock()
	defer q.mut.Unlock()

	return q.dequeueUnsafe()
}

func (q *linkedListQueue[T]) Dequeue(ctx context.Context) (T, error) {
	if ctx.Err() != nil {
		var zero T
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return zero, ErrTimeout
		}
		return zero, ErrCanceled
	}

	q.mut.Lock()
	defer q.mut.Unlock()

	stop := make(chan struct{})
	defer close(stop)

	go func() {
		select {
		case <-ctx.Done():
			q.mut.Lock()
			q.cond.Broadcast()
			q.mut.Unlock()
		case <-stop:
			return
		}
	}()

	for {
		if item, ok := q.dequeueUnsafe(); ok {
			return item, nil
		}

		if err := ctx.Err(); err != nil {
			var zero T
			if errors.Is(err, context.DeadlineExceeded) {
				return zero, ErrTimeout
			}
			return zero, ErrCanceled
		}

		q.cond.Wait()
	}
}

func (q *linkedListQueue[T]) Clear() {
	q.mut.Lock()
	defer q.mut.Unlock()

	atomic.StoreInt64(&q.size, 0)
	q.head = nil
	q.tail = nil

	q.cond.Broadcast()
}

func (q *linkedListQueue[T]) Size() int {
	return int(atomic.LoadInt64(&q.size))
}

func (q *linkedListQueue[T]) IsEmpty() bool {
	return atomic.LoadInt64(&q.size) == 0
}

func New[T any]() Queue[T] {
	var mut sync.Mutex
	return &linkedListQueue[T]{
		mut:  &mut,
		cond: sync.NewCond(&mut),
	}
}
