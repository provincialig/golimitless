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
			t.Fatal(size)
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
}
