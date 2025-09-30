package ds_test

import (
	"provincialig/golimitless/ds"
	"testing"
	"time"
)

func TestExpireSet(t *testing.T) {
	t.Run("Expire", func(t *testing.T) {
		es := ds.NewExpireSet[int]()
		es.Add(10, 100*time.Millisecond)

		time.Sleep(200 * time.Millisecond)

		if es.Has(10) {
			t.Fatal("Expire set should not have element 10")
		}
	})

	t.Run("NotExpire", func(t *testing.T) {
		es := ds.NewExpireSet[int]()
		es.Add(10, 2*time.Second)

		if !es.Has(10) {
			t.Fatal("Expire set should have element 10")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		es := ds.NewExpireSet[int]()
		es.Add(1, 100*time.Millisecond)
		es.Add(2, 200*time.Millisecond)

		time.Sleep(500 * time.Millisecond)

		if !es.IsEmpty() {
			t.Fatal("Expire must be empty")
		}
	})

	t.Run("NotEmpty", func(t *testing.T) {
		es := ds.NewExpireSet[int]()
		es.Add(1, 10*time.Second)
		es.Add(2, 20*time.Second)

		if es.IsEmpty() {
			t.Fatal("Expire must not be empty")
		}
	})
}
