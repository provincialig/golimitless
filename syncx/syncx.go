package syncx

import "sync"

func MutexBlockWithValue[T any](mut *sync.Mutex, fn func() (T, error)) (T, error) {
	mut.Lock()
	defer mut.Unlock()

	return fn()
}

func MutexBlock(mut *sync.Mutex, fn func() error) error {
	mut.Lock()
	defer mut.Unlock()

	return fn()
}
