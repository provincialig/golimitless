package slicesx_test

import (
	"reflect"
	"testing"

	"github.com/provincialig/golimitless/slicesx"
)

func Test_Filter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	out := slicesx.Filter(in, func(el int) bool { return el%2 == 0 })
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
	}

	empty := slicesx.Filter([]int{}, func(el int) bool { return true })
	if len(empty) != 0 {
		t.Fatalf("expected empty slice, got %v", empty)
	}

	none := slicesx.Filter(in, func(el int) bool { return false })
	if len(none) != 0 {
		t.Fatalf("expected empty slice, got %v", none)
	}
}

func Test_Map(t *testing.T) {
	in := []int{1, 2, 3}
	out := slicesx.Map(in, func(el int) int { return el * 2 })
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
	}

	out2 := slicesx.Map(in, func(el int) string { return string(rune('a' + el - 1)) })
	expected2 := []string{"a", "b", "c"}
	if !reflect.DeepEqual(out2, expected2) {
		t.Fatalf("expected %v, got %v", expected2, out2)
	}
}

func Test_Reduce(t *testing.T) {
	in := []int{1, 2, 3, 4}
	sum := slicesx.Reduce(in, 0, func(acc int, el int) int { return acc + el })
	if sum != 10 {
		t.Fatalf("expected 10, got %d", sum)
	}

	concat := slicesx.Reduce([]string{"go", "-", "limitless"}, "", func(acc string, el string) string { return acc + el })
	if concat != "go-limitless" {
		t.Fatalf("expected 'go-limitless', got %q", concat)
	}
}

func Test_ForEach(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	sum := 0
	count := 0
	slicesx.ForEach(in, func(el int) bool {
		if el >= 3 {
			return false
		}
		sum += el
		count++
		return true
	})
	if sum != 3 {
		t.Fatalf("expected sum 3 (1+2), got %d", sum)
	}
	if count != 2 {
		t.Fatalf("expected count 2, got %d", count)
	}
}

func TestMapToSlice_SliceToMap(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	s := slicesx.MapToSlice(m)
	back := slicesx.SliceToMap(s)
	if !reflect.DeepEqual(back, m) {
		t.Fatalf("expected %v, got %v", m, back)
	}
}
