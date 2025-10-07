package islice

import (
	"slices"
	"sync"
)

type ISlice[T comparable, K comparable] interface {
	Get(key T) ([]K, bool)
	Has(key T) bool
	Append(key T, value K)
	Contains(key T, value K) bool
	RemoveIndex(key T)
	RemoveElement(key T, index int)
	Range(fn func(key T, value []K) bool)
	IsEmpty(key T) bool
	Clear()
}

type myIndexedSlice[T comparable, K comparable] struct {
	m   map[T][]K
	mut sync.Mutex
}

func (is *myIndexedSlice[T, K]) Get(key T) ([]K, bool) {
	is.mut.Lock()
	defer is.mut.Unlock()

	v, ok := is.m[key]
	return v, ok
}

func (is *myIndexedSlice[T, K]) Has(key T) bool {
	is.mut.Lock()
	defer is.mut.Unlock()

	_, ok := is.m[key]
	return ok
}

func (is *myIndexedSlice[T, K]) RemoveIndex(key T) {
	is.mut.Lock()
	defer is.mut.Unlock()

	delete(is.m, key)
}

func (is *myIndexedSlice[T, K]) Append(key T, value K) {
	is.mut.Lock()
	defer is.mut.Unlock()

	_, ok := is.m[key]
	if !ok {
		is.m[key] = []K{}
	}

	is.m[key] = append(is.m[key], value)
}

func (is *myIndexedSlice[T, K]) Contains(key T, value K) bool {
	is.mut.Lock()
	defer is.mut.Unlock()

	v, ok := is.m[key]
	return ok && slices.Contains(v, value)
}

func (is *myIndexedSlice[T, K]) RemoveElement(key T, index int) {
	is.mut.Lock()
	defer is.mut.Unlock()

	_, ok := is.m[key]
	if ok && index >= 0 && index < len(is.m[key]) {
		is.m[key] = slices.Delete(is.m[key], index, index+1)
	}
}

func (is *myIndexedSlice[T, K]) Range(fn func(key T, value []K) bool) {
	for k, v := range is.m {
		if !fn(k, v) {
			return
		}
	}
}

func (is *myIndexedSlice[T, K]) IsEmpty(key T) bool {
	is.mut.Lock()
	defer is.mut.Unlock()

	v, ok := is.m[key]
	return ok && len(v) == 0
}

func (is *myIndexedSlice[T, K]) Clear() {
	is.mut.Lock()
	defer is.mut.Unlock()

	is.m = map[T][]K{}
}

func New[T comparable, K comparable]() ISlice[T, K] {
	return &myIndexedSlice[T, K]{m: map[T][]K{}}
}
