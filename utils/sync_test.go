package utils_test

import (
	"provincialig/golimitless/utils"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("MutexBlockWithValue", func(t *testing.T) {
		var mut sync.Mutex
		counter := 0
		const goroutines = 100
		wg := sync.WaitGroup{}
		wg.Add(goroutines)

		for range goroutines {
			go func() {
				defer wg.Done()
				_, err := utils.MutexBlockWithValue(&mut, func() (int, error) {
					tmp := counter
					tmp++
					counter = tmp
					return counter, nil
				})
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}()
		}
		wg.Wait()
		if counter != goroutines {
			t.Fatalf("expected counter to be %d, got %d", goroutines, counter)
		}
	})

	t.Run("MutexBlock", func(t *testing.T) {
		var mut sync.Mutex
		counter := 0
		const goroutines = 100
		wg := sync.WaitGroup{}
		wg.Add(goroutines)

		for range goroutines {
			go func() {
				defer wg.Done()
				err := utils.MutexBlock(&mut, func() error {
					tmp := counter
					tmp++
					counter = tmp
					return nil
				})
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}()
		}
		wg.Wait()
		if counter != goroutines {
			t.Fatalf("expected counter to be %d, got %d", goroutines, counter)
		}
	})
}
