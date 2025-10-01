package expireset_test

import (
	"testing"
	"time"

	"github.com/provincialig/golimitless/expireset"
)

func Test_Expire(t *testing.T) {
	es := expireset.New[int]()
	es.Add(10, 100*time.Millisecond)

	time.Sleep(200 * time.Millisecond)

	if es.Has(10) {
		t.Fatal("Expire set should not have element 10")
	}
}

func Test_NotExpire(t *testing.T) {
	es := expireset.New[int]()
	es.Add(10, 2*time.Second)

	if !es.Has(10) {
		t.Fatal("Expire set should have element 10")
	}
}

func Test_Empty(t *testing.T) {
	es := expireset.New[int]()
	es.Add(1, 100*time.Millisecond)
	es.Add(2, 200*time.Millisecond)

	time.Sleep(500 * time.Millisecond)

	if !es.IsEmpty() {
		t.Fatal("Expire must be empty")
	}
}

func Test_NotEmpty(t *testing.T) {
	es := expireset.New[int]()
	es.Add(1, 10*time.Second)
	es.Add(2, 20*time.Second)

	if es.IsEmpty() {
		t.Fatal("Expire must not be empty")
	}
}

func Test_Delete(t *testing.T) {
	es := expireset.New[int]()
	es.Add(42, 2*time.Second)
	if !es.Has(42) {
		t.Fatal("should have 42 before delete")
	}
	es.Delete(42)
	if es.Has(42) {
		t.Fatal("should not have 42 after delete")
	}
}

func Test_ExpireTime(t *testing.T) {
	es := expireset.New[int]()
	es.Add(7, 500*time.Millisecond)
	et, ok := es.ExpireTime(7)
	if !ok {
		t.Fatal("ExpireTime should return true for existing element")
	}
	if time.Until(et) <= 0 {
		t.Fatal("ExpireTime should be in the future")
	}
}

func Test_Iterator_Clear(t *testing.T) {
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
}
