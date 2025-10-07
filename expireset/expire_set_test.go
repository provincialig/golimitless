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

	if ok, _ := es.Has(10); ok {
		t.Fatal("Expire set should not have element 10")
	}
}

func Test_NotExpire(t *testing.T) {
	es := expireset.New[int]()
	es.Add(10, 2*time.Second)

	if ok, _ := es.Has(10); !ok {
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
	if ok, _ := es.Has(42); !ok {
		t.Fatal("should have 42 before delete")
	}
	es.Delete(42)
	if ok, _ := es.Has(42); ok {
		t.Fatal("should not have 42 after delete")
	}
}

func Test_ExpireTime(t *testing.T) {
	es := expireset.New[int]()
	es.Add(7, 500*time.Millisecond)
	ok, et := es.Has(7)
	if !ok {
		t.Fatal("ExpireTime should return true for existing element")
	}
	if time.Until(et) <= 0 {
		t.Fatal("ExpireTime should be in the future")
	}
}
