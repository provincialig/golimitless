package ds_test

import (
	"provincialig/golimitless/ds"
	"strings"
	"testing"
)

func TestIndexedSlice(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		s := ds.NewIndexedSlice[int, string]()

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
	})
}
