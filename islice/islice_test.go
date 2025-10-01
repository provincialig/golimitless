package islice_test

import (
	"strings"
	"testing"

	"github.com/provincialig/golimitless/islice"
)

func Test_Get(t *testing.T) {
	s := islice.New[int, string]()

	s.Append(1, "ciao")
	s.Append(1, " ")
	s.Append(1, "mondo")

	v, ok := s.Get(1)
	if !ok {
		t.Fatal("Slice must be contain index 1")
	}

	res := strings.Join(v, "")
	if res != "ciao mondo" {
		t.Fatalf("Result must be: ciao mondo: %s", res)
	}
}

func Test_Has(t *testing.T) {
	s := islice.New[int, int]()

	s.Append(1, 1)
	s.Append(2, 1)
	s.Append(1, 2)

	if !s.Has(1) {
		t.Fatal("Slice must be contain index 1")
	}

	if !s.Has(2) {
		t.Fatal("Slice must be contain index 2")
	}

	if s.Has(4) {
		t.Fatal("Slice must not be contain index 4")
	}
}

func Test_Remove(t *testing.T) {
	s := islice.New[int, int]()
	s.Append(1, 1)
	s.Append(1, 2)
	s.Append(1, 3)

	s.Remove(2, 6)
	s.Remove(1, 0)

	v, ok := s.Get(1)
	if !ok {
		t.Fatal("Slice must be contain index 1")
	}

	size := len(v)
	if size != 2 {
		t.Fatalf("Size of index 1 must be 2: %d", size)
	}
}

func Test_Contains(t *testing.T) {
	s := islice.New[int, int]()
	s.Append(1, 1)
	s.Append(1, 2)
	s.Append(1, 3)

	if !s.Contains(1, 2) {
		t.Fatal("Index 1 must be contains element 2")
	}

	if s.Contains(2, 2) {
		t.Fatal("Index 2 must be not contains element 2")
	}
}

func Test_IsEmpty(t *testing.T) {
	s := islice.New[int, int]()
	if s.IsEmpty(1) {
		t.Fatal("IsEmpty should be false for non-existent key (len undefined)")
	}
	s.Append(1, 10)
	if s.IsEmpty(1) {
		t.Fatal("IsEmpty should be false when one element exists")
	}
	s.Remove(1, 0)
	if !s.IsEmpty(1) {
		t.Fatal("IsEmpty should be true after removing the only element")
	}
}

func Test_DeleteKey(t *testing.T) {
	s := islice.New[int, int]()
	s.Append(1, 1)
	s.Append(1, 2)
	s.Append(2, 3)

	s.Delete(1)
	if s.Has(1) {
		t.Fatal("key 1 should be deleted")
	}
	if !s.Has(2) {
		t.Fatal("key 2 should still exist")
	}
}
