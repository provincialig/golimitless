package islice

import (
	"slices"

	"github.com/provincialig/golimitless/supermap"
)

type ISlice[T comparable, K comparable] interface {
	Get(key T) ([]K, bool)
	Has(key T) bool
	Delete(key T)
	Append(key T, value K)
	Contains(key T, value K) bool
	Remove(key T, index int)
	IsEmpty(key T) bool
}

type myIndexedSlice[T comparable, K comparable] struct {
	m supermap.SuperMap[T, *[]K]
}

func (is *myIndexedSlice[T, K]) Get(key T) ([]K, bool) {
	v, ok := is.m.Get(key)
	if !ok || v == nil {
		return nil, false
	}
	return *v, true
}

func (is *myIndexedSlice[T, K]) Has(key T) bool {
	return is.m.Has(key)
}

func (is *myIndexedSlice[T, K]) Delete(key T) {
	is.m.Delete(key)
}

func (is *myIndexedSlice[T, K]) Append(key T, value K) {
	if v, ok := is.m.Get(key); ok && v != nil {
		*v = append(*v, value)
		return
	}

	is.m.Set(key, &[]K{value})
}

func (is *myIndexedSlice[T, K]) Contains(key T, value K) bool {
	v, ok := is.m.Get(key)
	return ok && slices.Contains(*v, value)
}

func (is *myIndexedSlice[T, K]) Remove(key T, index int) {
	v, ok := is.m.Get(key)
	if ok && v != nil && index >= 0 && index < len(*v) {
		*v = append((*v)[:index], (*v)[index+1:]...)
	}
}

func (is *myIndexedSlice[T, K]) IsEmpty(key T) bool {
	v, ok := is.m.Get(key)
	return ok && len(*v) == 0
}

func New[T comparable, K comparable]() ISlice[T, K] {
	return &myIndexedSlice[T, K]{
		m: supermap.New[T, *[]K](),
	}
}
