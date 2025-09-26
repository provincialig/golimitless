package ds

import (
	"sync"
	"sync/atomic"
)

type SafeMapItem[T any, K any] struct {
	Key   T
	Value K
}

type SafeMap[T comparable, K any] interface {
	Get(key T) (K, bool)
	Set(key T, value K)
	Has(key T) bool
	Delete(key T)
	Clear()
	Range(fn func(key T, value K) bool)
	Size() int
	Keys() []T
	Values() []K
	ToSlice() []SafeMapItem[T, K]
}

type mySafeMap[T comparable, K any] struct {
	m     sync.Map
	count int64
}

func (sm *mySafeMap[T, K]) Get(key T) (K, bool) {
	v, ok := sm.m.Load(key)
	return v.(K), ok
}

func (sm *mySafeMap[T, K]) Set(key T, value K) {
	if _, ok := sm.m.Load(key); !ok {
		sm.m.Store(key, value)
		atomic.AddInt64(&sm.count, 1)
	}
}

func (sm *mySafeMap[T, K]) Has(key T) bool {
	_, ok := sm.m.Load(key)
	return ok
}

func (sm *mySafeMap[T, K]) Delete(key T) {
	if _, ok := sm.m.Load(key); ok {
		sm.m.Delete(key)
		atomic.AddInt64(&sm.count, -1)
	}
}

func (sm *mySafeMap[T, K]) Clear() {
	sm.m.Clear()
	atomic.StoreInt64(&sm.count, 0)
}

func (sm *mySafeMap[T, K]) Range(fn func(key T, value K) bool) {
	sm.m.Range(func(k, v any) bool {
		return fn(k.(T), v.(K))
	})
}

func (sm *mySafeMap[T, K]) Size() int {
	return int(atomic.LoadInt64(&sm.count))
}

func (sm *mySafeMap[T, K]) Keys() []T {
	res := []T{}

	sm.Range(func(key T, _ K) bool {
		res = append(res, key)
		return true
	})

	return res
}

func (sm *mySafeMap[T, K]) Values() []K {
	res := []K{}

	sm.Range(func(_ T, value K) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (sm *mySafeMap[T, K]) ToSlice() []SafeMapItem[T, K] {
	res := []SafeMapItem[T, K]{}

	sm.Range(func(key T, value K) bool {
		res = append(res, SafeMapItem[T, K]{
			Key:   key,
			Value: value,
		})
		return true
	})

	return res
}

func NewSafeMap[T comparable, K any]() SafeMap[T, K] {
	return &mySafeMap[T, K]{
		m:     sync.Map{},
		count: 0,
	}
}
