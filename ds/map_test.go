package ds_test

import (
	"provincialig/golimitless/ds"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		m := ds.NewMap[int, int]()
		m.Set(1, 2)
		m.Set(2, 2)
		m.Set(1, 3)
		size := m.Size()
		if size != 2 {
			t.Fatalf("Map size must be 2: %d", size)
		}
	})

	t.Run("Has", func(t *testing.T) {
		m := ds.NewMap[int, int]()
		m.Set(1, 1)
		if !m.Has(1) {
			t.Fatal("Map must be have 1")
		}
		if m.Has(2) {
			t.Fatal("Map must not be have 2")
		}
	})

	t.Run("Get", func(t *testing.T) {
		m := ds.NewMap[string, int]()
		m.Set("a", 10)
		val, ok := m.Get("a")
		if !ok || val != 10 {
			t.Fatalf("expected 10, got %v", val)
		}
		_, ok = m.Get("b")
		if ok {
			t.Fatal("expected false for non-existent key")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		m := ds.NewMap[int, string]()
		m.Set(1, "x")
		m.Delete(1)
		if m.Has(1) {
			t.Fatal("key should be deleted")
		}
		if m.Size() != 0 {
			t.Fatal("size should be 0 after delete")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		m := ds.NewMap[int, int]()
		m.Set(1, 1)
		m.Set(2, 2)
		m.Clear()
		if m.Size() != 0 {
			t.Fatal("map should be empty after clear")
		}
	})

	t.Run("Keys", func(t *testing.T) {
		m := ds.NewMap[string, int]()
		m.Set("a", 1)
		m.Set("b", 2)
		keys := m.Keys()
		if len(keys) != 2 {
			t.Fatal("keys length should be 2")
		}
	})

	t.Run("Values", func(t *testing.T) {
		m := ds.NewMap[int, string]()
		m.Set(1, "a")
		m.Set(2, "b")
		values := m.Values()
		if len(values) != 2 {
			t.Fatal("values length should be 2")
		}
	})

	t.Run("ToSlice", func(t *testing.T) {
		m := ds.NewMap[int, string]()
		m.Set(1, "a")
		m.Set(2, "b")
		slice := m.ToSlice()
		if len(slice) != 2 {
			t.Fatal("slice length should be 2")
		}
	})

	t.Run("Range", func(t *testing.T) {
		m := ds.NewMap[int, int]()
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
	})
}
