package ds

import "slices"

type IndexedSlice[T comparable, K comparable] interface {
	Get(key T) ([]K, bool)
	Has(key T) bool
	Delete(key T)
	Append(key T, value K)
	Contains(key T, value K) bool
	Remove(key T, index int)
	IsEmpty(key T) bool
}

type myIndexedSlice[T comparable, K comparable] struct {
	m Map[T, *[]K]
}

func (is *myIndexedSlice[T, K]) Get(key T) ([]K, bool) {
	v, ok := is.m.Get(key)
	return *v, ok
}

func (is *myIndexedSlice[T, K]) Has(key T) bool {
	return is.m.Has(key)
}

func (is *myIndexedSlice[T, K]) Delete(key T) {
	is.m.Delete(key)
}

func (is *myIndexedSlice[T, K]) Append(key T, value K) {
	if !is.m.Has(key) {
		is.m.Set(key, &[]K{})
	}

	v, _ := is.m.Get(key)
	*v = append(*v, value)
}

func (is *myIndexedSlice[T, K]) Contains(key T, value K) bool {
	v, ok := is.m.Get(key)
	return ok && slices.Contains(*v, value)
}

func (is *myIndexedSlice[T, K]) Remove(key T, index int) {
	v, ok := is.m.Get(key)
	if ok {
		*v = append((*v)[:index], (*v)[index+1:]...)
	}
}

func (is *myIndexedSlice[T, K]) IsEmpty(key T) bool {
	v, ok := is.m.Get(key)
	return ok && len(*v) == 0
}

func NewIndexedSlice[T comparable, K comparable]() IndexedSlice[T, K] {
	return &myIndexedSlice[T, K]{
		m: NewMap[T, *[]K](),
	}
}
