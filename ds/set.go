package ds

type Set[T comparable] interface {
	Add(values ...T)
	Remove(values ...T)
	Has(value T) bool
	Range(fn func(value T) bool)
	Union(set Set[T]) Set[T]
	Intersect(set Set[T]) Set[T]
	Difference(set Set[T]) Set[T]
	ToSlice() []T
	Size() int
}

type mySet[T comparable] struct {
	s Map[T, struct{}]
}

func (s *mySet[T]) Add(values ...T) {
	for _, value := range values {
		s.s.Set(value, struct{}{})
	}
}

func (s *mySet[T]) Remove(values ...T) {
	for _, value := range values {
		s.s.Delete(value)
	}
}

func (s *mySet[T]) Has(value T) bool {
	return s.s.Has(value)
}

func (s *mySet[T]) Range(fn func(value T) bool) {
	s.s.Range(func(key T, _ struct{}) bool {
		return fn(key)
	})
}

func (ss *mySet[T]) Union(s Set[T]) Set[T] {
	res := NewSet[T]()

	ss.Range(func(value T) bool {
		res.Add(value)
		return true
	})
	s.Range(func(value T) bool {
		res.Add(value)
		return true
	})

	return res
}

func (ss *mySet[T]) Intersect(s Set[T]) Set[T] {
	res := NewSet[T]()

	smaller, larger := Set[T](ss), s
	if s.Size() < ss.Size() {
		smaller, larger = s, ss
	}

	smaller.Range(func(value T) bool {
		if larger.Has(value) {
			res.Add(value)
		}
		return true
	})

	return res
}

func (ss *mySet[T]) Difference(s Set[T]) Set[T] {
	res := NewSet[T]()

	ss.Range(func(value T) bool {
		if !s.Has(value) {
			res.Add(value)
		}
		return true
	})

	return res
}

func (s *mySet[T]) ToSlice() []T {
	res := []T{}

	s.Range(func(value T) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (s *mySet[T]) Size() int {
	return s.s.Size()
}

func NewSet[T comparable]() Set[T] {
	return &mySet[T]{
		s: NewMap[T, struct{}](),
	}
}
