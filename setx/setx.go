package setx

import "github.com/provincialig/golimitless/mapx"

type SetX[T comparable] interface {
	Add(values ...T)
	Remove(values ...T)
	Has(value T) bool
	Range(fn func(value T) bool)
	Union(set SetX[T]) SetX[T]
	Intersect(set SetX[T]) SetX[T]
	Difference(set SetX[T]) SetX[T]
	ToSlice() []T
	Size() int
}

type mySetX[T comparable] struct {
	s mapx.MapX[T, struct{}]
}

func (sx *mySetX[T]) Add(values ...T) {
	for _, value := range values {
		sx.s.Set(value, struct{}{})
	}
}

func (sx *mySetX[T]) Remove(values ...T) {
	for _, value := range values {
		sx.s.Delete(value)
	}
}

func (sx *mySetX[T]) Has(value T) bool {
	return sx.s.Has(value)
}

func (sx *mySetX[T]) Range(fn func(value T) bool) {
	sx.s.Range(func(key T, _ struct{}) bool {
		return fn(key)
	})
}

func (sx *mySetX[T]) Union(s SetX[T]) SetX[T] {
	res := New[T]()

	sx.Range(func(value T) bool {
		res.Add(value)
		return true
	})
	s.Range(func(value T) bool {
		res.Add(value)
		return true
	})

	return res
}

func (sx *mySetX[T]) Intersect(s SetX[T]) SetX[T] {
	res := New[T]()

	smaller, larger := SetX[T](sx), s
	if s.Size() < sx.Size() {
		smaller, larger = s, sx
	}

	smaller.Range(func(value T) bool {
		if larger.Has(value) {
			res.Add(value)
		}
		return true
	})

	return res
}

func (sx *mySetX[T]) Difference(s SetX[T]) SetX[T] {
	res := New[T]()

	sx.Range(func(value T) bool {
		if !s.Has(value) {
			res.Add(value)
		}
		return true
	})

	return res
}

func (sx *mySetX[T]) ToSlice() []T {
	res := []T{}

	sx.Range(func(value T) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (sx *mySetX[T]) Size() int {
	return sx.s.Size()
}

func New[T comparable]() SetX[T] {
	return &mySetX[T]{
		s: mapx.New[T, struct{}](),
	}
}
