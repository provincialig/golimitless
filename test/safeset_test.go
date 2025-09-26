package test

import (
	"provincialig/golimitless/ds"
	"testing"
)

func TestSetSize(t *testing.T) {
	set := ds.NewSafeSet[int]()
	set.Add(1, 2, 3, 4)
	set.Remove(2, 1, 5)
	set.Add(6)

	size := set.Size()
	if size != 3 {
		t.Fatal(size)
	}
}

func TestSetUnion(t *testing.T) {
	a := ds.NewSafeSet[int]()
	a.Add(1, 2, 3, 4)

	b := ds.NewSafeSet[int]()
	b.Add(2, 3, 4, 5)

	sum := 0

	union := a.Union(b)
	union.Range(func(value int) bool {
		sum += value
		return true
	})

	if sum != 15 {
		t.Fatal(union.ToSlice())
	}
}

func TestSetIntersect(t *testing.T) {
	a := ds.NewSafeSet[int]()
	a.Add(1, 2, 3, 4)

	b := ds.NewSafeSet[int]()
	b.Add(2, 3, 4, 5)

	sum := 0

	intersect := a.Intersect(b)
	intersect.Range(func(value int) bool {
		sum += value
		return true
	})

	if sum != 9 {
		t.Fatal(intersect.ToSlice())
	}
}

func TestSetDifference(t *testing.T) {
	a := ds.NewSafeSet[int]()
	a.Add(1, 2, 3, 4)

	b := ds.NewSafeSet[int]()
	b.Add(2, 3, 4, 5)

	sum := 0

	difference := a.Difference(b)
	difference.Range(func(value int) bool {
		sum += value
		return true
	})

	if sum != 1 {
		t.Fatal(difference.ToSlice())
	}
}
