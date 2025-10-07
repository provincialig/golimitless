package expireset

import (
	"sync"
	"time"
)

type ExpireSet[T comparable] interface {
	Add(value T, retain time.Duration)
	Has(value T) (bool, time.Time)
	Delete(value T)
	Clear()
	Size() int
	IsEmpty() bool
}

type myExpireSet[T comparable] struct {
	m   map[T]time.Time
	mut sync.Mutex
}

func (es *myExpireSet[T]) getUnsafe(value T) (time.Time, bool) {
	var zero time.Time

	v, ok := es.m[value]
	if !ok {
		return zero, false
	}

	if time.Now().After(v) {
		delete(es.m, value)
		return zero, false
	}

	return v, true
}

func (es *myExpireSet[T]) Add(value T, retain time.Duration) {
	es.mut.Lock()
	defer es.mut.Unlock()

	es.m[value] = time.Now().Add(retain)
}

func (es *myExpireSet[T]) Has(value T) (bool, time.Time) {
	es.mut.Lock()
	defer es.mut.Unlock()

	v, ok := es.getUnsafe(value)
	return ok, v
}

func (es *myExpireSet[T]) Delete(value T) {
	es.mut.Lock()
	defer es.mut.Unlock()

	delete(es.m, value)
}

func (es *myExpireSet[T]) Clear() {
	es.mut.Lock()
	defer es.mut.Unlock()

	es.m = map[T]time.Time{}
}

func (es *myExpireSet[T]) Size() int {
	es.mut.Lock()
	defer es.mut.Unlock()

	size := 0

	for k := range es.m {
		if _, ok := es.getUnsafe(k); ok {
			size++
		}
	}

	return size
}

func (es *myExpireSet[T]) IsEmpty() bool {
	es.mut.Lock()
	defer es.mut.Unlock()

	isEmpty := true

	for k := range es.m {
		if _, ok := es.getUnsafe(k); ok {
			isEmpty = false
			break
		}
	}

	return isEmpty
}

func New[T comparable]() ExpireSet[T] {
	return &myExpireSet[T]{m: map[T]time.Time{}}
}
