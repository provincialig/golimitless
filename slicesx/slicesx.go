package slicesx

func Filter[T any](silce []T, fn func(el T) bool) []T {
	result := []T{}

	for _, el := range silce {
		if fn(el) {
			result = append(result, el)
		}
	}

	return result
}

func Map[T any, K any](slice []T, fn func(el T) K) []K {
	result := []K{}

	for _, el := range slice {
		result = append(result, fn(el))
	}

	return result
}

func Reduce[T any, K any](slice []T, initial K, fn func(acc K, el T) K) K {
	acc := initial

	for _, el := range slice {
		acc = fn(acc, el)
	}

	return acc
}

func ForEach[T any](slice []T, fn func(el T) bool) {
	for _, el := range slice {
		if !fn(el) {
			return
		}
	}
}

type SliceItem[T any, K any] struct {
	Index T
	Value K
}

func MapToSlice[T comparable, K any](m map[T]K) []SliceItem[T, K] {
	res := []SliceItem[T, K]{}

	for k, v := range m {
		res = append(res, SliceItem[T, K]{
			Index: k,
			Value: v,
		})
	}

	return res
}

func SliceToMap[T comparable, K any](slice []SliceItem[T, K]) map[T]K {
	res := map[T]K{}

	for _, el := range slice {
		res[el.Index] = el.Value
	}

	return res
}
