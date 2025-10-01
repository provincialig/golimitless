package mapx

import (
	"sync"
	"sync/atomic"
)

type MapXItem[T any, K any] struct {
	Key   T
	Value K
}

type MapX[T comparable, K any] interface {
	Get(key T) (K, bool)
	Set(key T, value K)
	LoadOrStore(key T, value K) (K, bool)
	Has(key T) bool
	Delete(key T)
	Clear()
	Range(fn func(key T, value K) bool)
	Size() int
	Keys() []T
	Values() []K
	ToSlice() []MapXItem[T, K]
}

type myMapX[T comparable, K any] struct {
	m     sync.Map
	count int64
}

func (mx *myMapX[T, K]) Get(key T) (K, bool) {
	v, ok := mx.m.Load(key)
	if !ok {
		var zero K
		return zero, false
	}

	return v.(K), true
}

func (mx *myMapX[T, K]) Set(key T, value K) {
	if _, loaded := mx.m.LoadOrStore(key, value); !loaded {
		atomic.AddInt64(&mx.count, 1)
		return
	}
	mx.m.Store(key, value)
}

func (mx *myMapX[T, K]) LoadOrStore(key T, value K) (K, bool) {
	actual, loaded := mx.m.LoadOrStore(key, value)
	if !loaded {
		atomic.AddInt64(&mx.count, 1)
		return value, false
	}

	return actual.(K), true
}

func (mx *myMapX[T, K]) Has(key T) bool {
	_, ok := mx.m.Load(key)
	return ok
}

func (mx *myMapX[T, K]) Delete(key T) {
	if _, ok := mx.m.Load(key); ok {
		mx.m.Delete(key)
		atomic.AddInt64(&mx.count, -1)
	}
}

func (mx *myMapX[T, K]) Clear() {
	mx.m.Clear()
	atomic.StoreInt64(&mx.count, 0)
}

func (mx *myMapX[T, K]) Range(fn func(key T, value K) bool) {
	mx.m.Range(func(k, v any) bool {
		return fn(k.(T), v.(K))
	})
}

func (mx *myMapX[T, K]) Size() int {
	return int(atomic.LoadInt64(&mx.count))
}

func (mx *myMapX[T, K]) Keys() []T {
	res := []T{}

	mx.Range(func(key T, _ K) bool {
		res = append(res, key)
		return true
	})

	return res
}

func (mx *myMapX[T, K]) Values() []K {
	res := []K{}

	mx.Range(func(_ T, value K) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (mx *myMapX[T, K]) ToSlice() []MapXItem[T, K] {
	res := []MapXItem[T, K]{}

	mx.Range(func(key T, value K) bool {
		res = append(res, MapXItem[T, K]{
			Key:   key,
			Value: value,
		})
		return true
	})

	return res
}

func New[T comparable, K any]() MapX[T, K] {
	return &myMapX[T, K]{
		m:     sync.Map{},
		count: 0,
	}
}
