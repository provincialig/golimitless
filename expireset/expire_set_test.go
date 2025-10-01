package expireset_test

import (
	"testing"
	"time"

	"github.com/provincialig/golimitless/expireset"
)

func TestExpireSet(t *testing.T) {
	t.Run("Expire", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(10, 100*time.Millisecond)

		time.Sleep(200 * time.Millisecond)

		if es.Has(10) {
			t.Fatal("Expire set should not have element 10")
		}
	})

	t.Run("NotExpire", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(10, 2*time.Second)

		if !es.Has(10) {
			t.Fatal("Expire set should have element 10")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(1, 100*time.Millisecond)
		es.Add(2, 200*time.Millisecond)

		time.Sleep(500 * time.Millisecond)

		if !es.IsEmpty() {
			t.Fatal("Expire must be empty")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(1, 10*time.Second)
		es.Add(2, 20*time.Second)

		if es.IsEmpty() {
			t.Fatal("Expire must not be empty")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(42, 2*time.Second)
		if !es.Has(42) {
			t.Fatal("should have 42 before delete")
		}
		es.Delete(42)
		if es.Has(42) {
			t.Fatal("should not have 42 after delete")
		}
	})

	t.Run("ExpireTime", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(7, 500*time.Millisecond)
		et, ok := es.ExpireTime(7)
		if !ok {
			t.Fatal("ExpireTime should return true for existing element")
		}
		if time.Until(et) <= 0 {
			t.Fatal("ExpireTime should be in the future")
		}
	})

	t.Run("Iterator & Clear", func(t *testing.T) {
		es := expireset.New[int]()
		es.Add(1, 5*time.Second)
		es.Add(2, 5*time.Second)
		es.Add(3, 5*time.Second)

		sum := 0
		es.Iterator(func(value int) bool {
			sum += value
			return true
		})
		if sum != 6 {
			t.Fatalf("sum should be 6, got %d", sum)
		}

		es.Clear()
		if !es.IsEmpty() {
			t.Fatal("ExpireSet should be empty after Clear")
		}
	})
}
