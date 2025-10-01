package expireset

import (
	"sync"
	"time"
)

type ExpireSet[T comparable] interface {
	Add(value T, retain time.Duration)
	Has(value T) bool
	Delete(value T)
	ExpireTime(value T) (time.Time, bool)
	Iterator(fn func(value T) bool)
	Clear()
	IsEmpty() bool
}

type myExpireSet[T comparable] struct {
	m   map[T]time.Time
	mut *sync.Mutex
}

func (es *myExpireSet[T]) get(value T) (time.Time, bool) {
	v, ok := es.m[value]
	if !ok {
		return time.Time{}, false
	}

	if time.Now().After(v) {
		delete(es.m, value)
		return time.Time{}, false
	}

	return v, true
}

func (es *myExpireSet[T]) Add(value T, retain time.Duration) {
	es.mut.Lock()
	defer es.mut.Unlock()

	es.m[value] = time.Now().Add(retain)
}

func (es *myExpireSet[T]) Has(value T) bool {
	es.mut.Lock()
	defer es.mut.Unlock()

	_, ok := es.get(value)
	return ok
}

func (es *myExpireSet[T]) Delete(value T) {
	es.mut.Lock()
	defer es.mut.Unlock()

	delete(es.m, value)
}

func (es *myExpireSet[T]) ExpireTime(value T) (time.Time, bool) {
	es.mut.Lock()
	defer es.mut.Unlock()

	v, ok := es.get(value)
	return v, ok
}

func (es *myExpireSet[T]) Clear() {
	es.mut.Lock()
	defer es.mut.Unlock()

	es.m = map[T]time.Time{}
}

func (es *myExpireSet[T]) Iterator(fn func(value T) bool) {
	es.mut.Lock()
	defer es.mut.Unlock()

	for k := range es.m {
		if _, ok := es.get(k); ok {
			if !fn(k) {
				return
			}
		}
	}
}

func (es *myExpireSet[T]) IsEmpty() bool {
	found := false
	es.Iterator(func(value T) bool {
		found = true
		return false
	})
	return !found
}

func New[T comparable]() ExpireSet[T] {
	return &myExpireSet[T]{
		m:   map[T]time.Time{},
		mut: &sync.Mutex{},
	}
}
