package mapx_test

import (
	"sync"
	"testing"

	"github.com/provincialig/golimitless/mapx"
)

func Test_Size(t *testing.T) {
	m := mapx.New[int, int]()
	m.Set(1, 2)
	m.Set(2, 2)
	m.Set(1, 3)
	size := m.Size()
	if size != 2 {
		t.Fatalf("Map size must be 2: %d", size)
	}
}

func Test_Has(t *testing.T) {
	m := mapx.New[int, int]()
	m.Set(1, 1)
	if !m.Has(1) {
		t.Fatal("Map must be have 1")
	}
	if m.Has(2) {
		t.Fatal("Map must not be have 2")
	}
}

func Test_Get(t *testing.T) {
	m := mapx.New[string, int]()
	m.Set("a", 10)
	val, ok := m.Get("a")
	if !ok || val != 10 {
		t.Fatalf("expected 10, got %v", val)
	}
	_, ok = m.Get("b")
	if ok {
		t.Fatal("expected false for non-existent key")
	}
}

func Test_Delete(t *testing.T) {
	m := mapx.New[int, string]()
	m.Set(1, "x")
	m.Delete(1)
	if m.Has(1) {
		t.Fatal("key should be deleted")
	}
	if m.Size() != 0 {
		t.Fatal("size should be 0 after delete")
	}
}

func Test_Clear(t *testing.T) {
	m := mapx.New[int, int]()
	m.Set(1, 1)
	m.Set(2, 2)
	m.Clear()
	if m.Size() != 0 {
		t.Fatal("map should be empty after clear")
	}
}

func Test_Keys(t *testing.T) {
	m := mapx.New[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)
	keys := m.Keys()
	if len(keys) != 2 {
		t.Fatal("keys length should be 2")
	}
}

func Test_Values(t *testing.T) {
	m := mapx.New[int, string]()
	m.Set(1, "a")
	m.Set(2, "b")
	values := m.Values()
	if len(values) != 2 {
		t.Fatal("values length should be 2")
	}
}

func Test_ToSlice(t *testing.T) {
	m := mapx.New[int, string]()
	m.Set(1, "a")
	m.Set(2, "b")
	slice := m.ToSlice()
	if len(slice) != 2 {
		t.Fatal("slice length should be 2")
	}
}

func Test_Range(t *testing.T) {
	m := mapx.New[int, int]()
	m.Set(1, 10)
	m.Set(2, 20)
	sum := 0
	m.Range(func(k, v int) bool {
		sum += v
		return true
	})
	if sum != 30 {
		t.Fatalf("expected sum 30, got %d", sum)
	}
}

func Test_ConcurrentSetSameKey(t *testing.T) {
	m := mapx.New[int, int]()

	const goroutines = 1000
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(v int) {
			defer wg.Done()
			m.Set(1, v)
		}(i)
	}
	wg.Wait()

	if m.Size() != 1 {
		t.Fatalf("size should be 1 after concurrent sets on same key, got %d", m.Size())
	}

	val, ok := m.Get(1)
	if !ok {
		t.Fatal("key 1 should exist after concurrent sets")
	}
	if val < 0 || val >= goroutines {
		t.Fatalf("value should be one of the written values, got %d", val)
	}
}

func Test_OverwriteAfterConcurrentSets(t *testing.T) {
	m := mapx.New[int, int]()

	const goroutines = 200
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(v int) {
			defer wg.Done()
			m.Set(1, v)
		}(i)
	}
	wg.Wait()

	m.Set(1, 999)

	val, ok := m.Get(1)
	if !ok {
		t.Fatal("key 1 should exist")
	}
	if val != 999 {
		t.Fatalf("expected final value 999, got %d", val)
	}
	if m.Size() != 1 {
		t.Fatalf("size should remain 1, got %d", m.Size())
	}
}
