package retainer

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/provincialig/golimitless/queue"
)

type CancelFunc = func()

type Retainer[T comparable] interface {
	Add(data T, retain time.Duration)
	Get() (<-chan T, CancelFunc)
	Clean()
	Destroy()
}

type myRetainer[T comparable] struct {
	start bool

	out queue.Queue[T]
	m   map[T]time.Time

	mut sync.Mutex
	t   *time.Ticker

	ctx    context.Context
	cancel context.CancelFunc
}

func (r *myRetainer[T]) worker() {
	for range r.t.C {
		now := time.Now()

		r.mut.Lock()

		for k, v := range r.m {
			if now.After(v) {
				delete(r.m, k)
				r.out.Enqueue(k)
			}
		}

		r.mut.Unlock()
	}
}

func (r *myRetainer[T]) Add(data T, retain time.Duration) {
	r.mut.Lock()
	defer r.mut.Unlock()

	if !r.start {
		return
	}

	r.m[data] = time.Now().Add(retain)
}

func (r *myRetainer[T]) Get() (<-chan T, CancelFunc) {
	if !r.start {
		return nil, nil
	}

	ctx, cancel := context.WithCancel(r.ctx)

	out := make(chan T)

	go func() {
		defer close(out)

		for {
			data, err := r.out.Dequeue(ctx)
			if errors.Is(err, queue.ErrCanceled) {
				return
			}

			out <- data
		}
	}()

	return out, cancel
}

func (r *myRetainer[T]) Clean() {
	r.mut.Lock()
	defer r.mut.Unlock()

	if !r.start {
		return
	}

	r.m = map[T]time.Time{}
}

func (r *myRetainer[T]) Destroy() {
	r.mut.Lock()
	defer r.mut.Unlock()

	if !r.start {
		return
	}

	r.t.Stop()

	r.cancel()

	r.m = map[T]time.Time{}
	r.out.Clear()
	r.start = false
}

func New[T comparable]() Retainer[T] {
	ctx, cancel := context.WithCancel(context.Background())

	retainer := &myRetainer[T]{
		start:  true,
		out:    queue.New[T](),
		m:      map[T]time.Time{},
		t:      time.NewTicker(100 * time.Millisecond),
		ctx:    ctx,
		cancel: cancel,
	}
	go retainer.worker()

	return retainer
}
