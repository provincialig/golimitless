package test

import (
	"provincialig/golimitless/ds"
	"testing"
)

func TestMapSize(t *testing.T) {
	m := ds.NewSafeMap[int, int]()
	m.Set(1, 2)
	m.Set(2, 2)
	m.Set(1, 3)

	size := m.Size()
	if size != 2 {
		t.Fatal(size)
	}
}
