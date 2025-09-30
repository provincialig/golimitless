package ds_test

import (
	"log"
	"provincialig/golimitless/ds"
	"testing"
)

func TestDoubleMap(t *testing.T) {
	t.Run("Size", func(t *testing.T) {
		dm := ds.NewDoubleMap[int, int, string]()

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
		dm := ds.NewDoubleMap[int, int, string]()

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
		dm := ds.NewDoubleMap[string, string, int]()

		dm.Set("a", "b", 1)

		v, ok := dm.Get("a", "b")

		if !ok {
			t.Fatal("Map must be have a-b element")
		}

		if v != 1 {
			t.Fatal("Map a-b element must be 1")
		}

	})
}
