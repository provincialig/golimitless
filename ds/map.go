package ds

import (
	"sync"
	"sync/atomic"
)

type MapItem[T any, K any] struct {
	Key   T
	Value K
}

type Map[T comparable, K any] interface {
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

type myMap[T comparable, K any] struct {
	m     sync.Map
	count int64
}

func (m *myMap[T, K]) Get(key T) (K, bool) {
	v, ok := m.m.Load(key)
	return v.(K), ok
}

func (m *myMap[T, K]) Set(key T, value K) {
	if _, ok := m.m.Load(key); !ok {
		m.m.Store(key, value)
		atomic.AddInt64(&m.count, 1)
	}
}

func (m *myMap[T, K]) Has(key T) bool {
	_, ok := m.m.Load(key)
	return ok
}

func (m *myMap[T, K]) Delete(key T) {
	if _, ok := m.m.Load(key); ok {
		m.m.Delete(key)
		atomic.AddInt64(&m.count, -1)
	}
}

func (m *myMap[T, K]) Clear() {
	m.m.Clear()
	atomic.StoreInt64(&m.count, 0)
}

func (m *myMap[T, K]) Range(fn func(key T, value K) bool) {
	m.m.Range(func(k, v any) bool {
		return fn(k.(T), v.(K))
	})
}

func (m *myMap[T, K]) Size() int {
	return int(atomic.LoadInt64(&m.count))
}

func (m *myMap[T, K]) Keys() []T {
	res := []T{}

	m.Range(func(key T, _ K) bool {
		res = append(res, key)
		return true
	})

	return res
}

func (m *myMap[T, K]) Values() []K {
	res := []K{}

	m.Range(func(_ T, value K) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (m *myMap[T, K]) ToSlice() []MapItem[T, K] {
	res := []MapItem[T, K]{}

	m.Range(func(key T, value K) bool {
		res = append(res, MapItem[T, K]{
			Key:   key,
			Value: value,
		})
		return true
	})

	return res
}

func NewMap[T comparable, K any]() Map[T, K] {
	return &myMap[T, K]{
		m:     sync.Map{},
		count: 0,
	}
}
