package doublemap_test

import (
	"log"
	"testing"

	"github.com/provincialig/golimitless/doublemap"
)

func TestDoubleMap(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		dm := doublemap.New[int, int, string]()

		dm.Set(1, 1, "a")
		dm.Set(1, 2, "b")
		dm.Set(2, 1, "a")

		if size := dm.SizeRoot(); size != 2 {
			log.Fatalf("Root size must be 2: %d", size)
		}

		if size, _ := dm.SizeChild(1); size != 2 {
			log.Fatalf("Child 1 size must be 2: %d", size)
		}

		if size, _ := dm.SizeChild(2); size != 1 {
			log.Fatalf("Child 2 size must be 1: %d", size)
		}
	})

	t.Run("Has", func(t *testing.T) {
		dm := doublemap.New[int, int, string]()

		dm.Set(1, 1, "a")
		dm.Set(1, 2, "b")
		dm.Set(2, 1, "a")

		if !dm.Has(1, 1) {
			t.Fatal("Map must be have 1-1 element")
		}

		if !dm.Has(2, 1) {
			t.Fatal("Map must be have 2-1 element")
		}

		if dm.Has(2, 2) {
			t.Fatal("Map must not be have 2-2 element")
		}

		if dm.Has(3, 6) {
			t.Fatal("Map must not be have 3-6 element")
		}
	})

	t.Run("Get", func(t *testing.T) {
		dm := doublemap.New[string, string, int]()

		dm.Set("a", "b", 1)

		v, ok := dm.Get("a", "b")

		if !ok {
			t.Fatal("Map must be have a-b element")
		}

		if v != 1 {
			t.Fatal("Map a-b element must be 1")
		}

	})

	t.Run("Delete & RootKeys & ChildKeys", func(t *testing.T) {
		dm := doublemap.New[int, int, string]()
		dm.Set(1, 1, "a")
		dm.Set(1, 2, "b")
		dm.Set(2, 1, "c")

		keys := dm.RootKeys()
		if len(keys) != 2 {
			t.Fatalf("root keys length should be 2, got %d", len(keys))
		}

		child1, ok := dm.ChildKeys(1)
		if !ok || len(child1) != 2 {
			t.Fatal("child keys for 1 should exist and have length 2")
		}

		dm.Delete(1, 1)
		if dm.Has(1, 1) {
			t.Fatal("should not have 1-1 after delete")
		}

		child1, ok = dm.ChildKeys(1)
		if !ok || len(child1) != 1 {
			t.Fatalf("child 1 length should be 1, got %d", len(child1))
		}
	})

	t.Run("ClearRoot & ClearChild", func(t *testing.T) {
		dm := doublemap.New[int, int, int]()
		dm.Set(1, 1, 1)
		dm.Set(1, 2, 2)
		dm.Set(2, 1, 3)

		dm.ClearChild(1)
		if size, ok := dm.SizeChild(1); !ok || size != 0 {
			t.Fatalf("child 1 size should be 0, got %d (ok=%v)", size, ok)
		}

		dm.ClearRoot()
		if dm.SizeRoot() != 0 {
			t.Fatal("root should be empty after ClearRoot")
		}
	})
}
