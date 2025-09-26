package ds

type SafeSet[T comparable] interface {
	Add(values ...T)
	Remove(values ...T)
	Has(value T) bool
	Range(fn func(value T) bool)
	Union(set SafeSet[T]) SafeSet[T]
	Intersect(set SafeSet[T]) SafeSet[T]
	Difference(set SafeSet[T]) SafeSet[T]
	ToSlice() []T
	Size() int
}

type mySafeSet[T comparable] struct {
	s SafeMap[T, struct{}]
}

func (ss *mySafeSet[T]) Add(values ...T) {
	for _, value := range values {
		ss.s.Set(value, struct{}{})
	}
}

func (ss *mySafeSet[T]) Remove(values ...T) {
	for _, value := range values {
		ss.s.Delete(value)
	}
}

func (ss *mySafeSet[T]) Has(value T) bool {
	return ss.s.Has(value)
}

func (ss *mySafeSet[T]) Range(fn func(value T) bool) {
	ss.s.Range(func(key T, _ struct{}) bool {
		return fn(key)
	})
}

func (ss *mySafeSet[T]) Union(s SafeSet[T]) SafeSet[T] {
	res := NewSafeSet[T]()

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

func (ss *mySafeSet[T]) Intersect(s SafeSet[T]) SafeSet[T] {
	res := NewSafeSet[T]()

	smaller, larger := SafeSet[T](ss), s
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

func (ss *mySafeSet[T]) Difference(s SafeSet[T]) SafeSet[T] {
	res := NewSafeSet[T]()

	ss.Range(func(value T) bool {
		if !s.Has(value) {
			res.Add(value)
		}
		return true
	})

	return res
}

func (ss *mySafeSet[T]) ToSlice() []T {
	res := []T{}

	ss.Range(func(value T) bool {
		res = append(res, value)
		return true
	})

	return res
}

func (ss *mySafeSet[T]) Size() int {
	return ss.s.Size()
}

func NewSafeSet[T comparable]() SafeSet[T] {
	return &mySafeSet[T]{
		s: NewSafeMap[T, struct{}](),
	}
}
