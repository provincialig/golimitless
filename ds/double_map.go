package ds

type DoubleMap[T comparable, K comparable, V any] interface {
	Set(key1 T, key2 K, value V)
	Get(key1 T, key2 K) (V, bool)
	Has(key1 T, key2 K) bool
	Delete(key1 T, key2 K)
	RootKeys() []T
	ChildKeys(key T) ([]K, bool)
	SizeRoot() int
	SizeChild(key T) (int, bool)
	ClearRoot()
	ClearChild(key T)
}

type myDoubleMap[T comparable, K comparable, V any] struct {
	m Map[T, Map[K, V]]
}

func (dm *myDoubleMap[T, K, V]) Set(key1 T, key2 K, value V) {
	if !dm.m.Has(key1) {
		dm.m.Set(key1, NewMap[K, V]())
	}

	child, _ := dm.m.Get(key1)
	child.Set(key2, value)
}

func (dm *myDoubleMap[T, K, V]) Get(key1 T, key2 K) (V, bool) {
	child, ok := dm.m.Get(key1)
	if !ok {
		var zero V
		return zero, false
	}

	v, ok := child.Get(key2)
	return v, ok
}

func (dm *myDoubleMap[T, K, V]) Has(key1 T, key2 K) bool {
	child, ok := dm.m.Get(key1)
	if !ok {
		return false
	}

	_, ok = child.Get(key2)
	return ok
}

func (dm *myDoubleMap[T, K, V]) Delete(key1 T, key2 K) {
	child, ok := dm.m.Get(key1)
	if ok {
		child.Delete(key2)
	}
}

func (dm *myDoubleMap[T, K, V]) RootKeys() []T {
	return dm.m.Keys()
}

func (dm *myDoubleMap[T, K, V]) ChildKeys(key T) ([]K, bool) {
	child, ok := dm.m.Get(key)
	if !ok {
		return nil, false
	}
	return child.Keys(), true
}

func (dm *myDoubleMap[T, K, V]) SizeRoot() int {
	return dm.m.Size()
}

func (dm *myDoubleMap[T, K, V]) SizeChild(key T) (int, bool) {
	child, ok := dm.m.Get(key)
	if !ok {
		return 0, false
	}
	return child.Size(), true
}

func (dm *myDoubleMap[T, K, V]) ClearRoot() {
	dm.m.Clear()
}

func (dm *myDoubleMap[T, K, V]) ClearChild(key T) {
	child, ok := dm.m.Get(key)
	if ok {
		child.Clear()
	}
}

func NewDoubleMap[T comparable, K comparable, V any]() DoubleMap[T, K, V] {
	return &myDoubleMap[T, K, V]{
		m: NewMap[T, Map[K, V]](),
	}
}
