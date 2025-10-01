package supermap

import (
	"sync"
	"sync/atomic"
)

type MapItem[T any, K any] struct {
	Key   T
	Value K
}

type SuperMap[T comparable, K any] interface {
	Get(key T) (K, bool)
	Set(key T, value K)
	Has(key T) bool
	Delete(key T)
	Clear()
	Range(fn func(key T, value K) bool)
	Size() int
	Keys() []T
	Values() []K
	ToSlice() []MapItem[T, K]
}

type mySuperMap[T comparable, K any] struct {
	m     sync.Map
	count int64
}

func (sm *mySuperMap[T, K]) Get(key T) (K, bool) {
	v, ok := sm.m.Load(key)
	if !ok {
		var zero K
		return zero, false
	}

	return v.(K), true
}

func (sm *mySuperMap[T, K]) Set(key T, value K) {
	if _, loaded := sm.m.LoadOrStore(key, value); !loaded {
		atomic.AddInt64(&sm.count, 1)
		return
	}
	sm.m.Store(key, value)
}

func (sm *mySuperMap[T, K]) Has(key T) bool {
	_, ok := sm.m.Load(key)
	return ok
}

func (sm *mySuperMap[T, K]) Delete(key T) {
	if _, ok := sm.m.Load(key); ok {
		sm.m.Delete(key)
		atomic.AddInt64(&sm.count, -1)
	}
}

func (sm *mySuperMap[T, K]) Clear() {
	sm.m.Clear()
	atomic.StoreInt64(&sm.count, 0)
}

func (sm *mySuperMap[T, K]) Range(fn func(key T, value K) bool) {
	sm.m.Range(func(k, v any) bool {
		return fn(k.(T), v.(K))
	})
}

func (sm *mySuperMap[T, K]) Size() int {
	return int(atomic.LoadInt64(&sm.count))
}

func (sm *mySuperMap[T, K]) Keys() []T {
	res := []T{}

	sm.Range(func(key T, _ K) bool {
		res = append(res, key)
		return true
	})

	return res
}

func (sm *mySuperMap[T, K]) Values() []K {
	res := []K{}

	sm.Range(func(_ T, value K) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (sm *mySuperMap[T, K]) ToSlice() []MapItem[T, K] {
	res := []MapItem[T, K]{}

	sm.Range(func(key T, value K) bool {
		res = append(res, MapItem[T, K]{
			Key:   key,
			Value: value,
		})
		return true
	})

	return res
}

func New[T comparable, K any]() SuperMap[T, K] {
	return &mySuperMap[T, K]{
		m:     sync.Map{},
		count: 0,
	}
}
