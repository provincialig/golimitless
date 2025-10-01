package set_test

import (
	"testing"

	"github.com/provincialig/golimitless/set"
)

func Test_Size(t *testing.T) {
	set := set.New[int]()
	set.Add(1, 2, 3, 4)
	set.Remove(2, 1, 5)
	set.Add(6)

	size := set.Size()
	if size != 3 {
		t.Fatal(size)
	}
}

func Test_Union(t *testing.T) {
	a := set.New[int]()
	a.Add(1, 2, 3, 4)

	b := set.New[int]()
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

func Test_Intersection(t *testing.T) {
	a := set.New[int]()
	a.Add(1, 2, 3, 4)

	b := set.New[int]()
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

func Test_Difference(t *testing.T) {
	a := set.New[int]()
	a.Add(1, 2, 3, 4)

	b := set.New[int]()
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

func Test_Has_Remove(t *testing.T) {
	set := set.New[int]()
	set.Add(1, 2)
	if !set.Has(1) || !set.Has(2) {
		t.Fatal("set should contain 1 and 2")
	}
	set.Remove(2)
	if set.Has(2) {
		t.Fatal("set should not contain 2 after remove")
	}
}

func Test_Range(t *testing.T) {
	set := set.New[int]()
	set.Add(1, 2, 3)
	sum := 0
	set.Range(func(value int) bool {
		sum += value
		return true
	})
	if sum != 6 {
		t.Fatalf("expected sum 6, got %d", sum)
	}
}

func Test_ToSlice(t *testing.T) {
	set := set.New[string]()
	set.Add("a", "b")
	slice := set.ToSlice()
	if len(slice) != 2 {
		t.Fatalf("slice length should be 2, got %d", len(slice))
	}
}
