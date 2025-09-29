package utils

func SliceFilter[T any](silce []T, fn func(el T) bool) []T {
	result := []T{}

	for _, el := range silce {
		if fn(el) {
			result = append(result, el)
		}
	}

	return result
}

func SliceMap[T any, K any](slice []T, fn func(el T) (K, error)) ([]K, error) {
	result := []K{}

	for _, el := range slice {
		r, err := fn(el)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func SliceForEach[T any](slice []T, fn func(el T) bool) {
	for _, el := range slice {
		if !fn(el) {
			return
		}
	}
}

func SliceReduce[T any, K any](slice []T, initial K, fn func(acc K, el T) (K, error)) (K, error) {
	acc := initial

	for _, el := range slice {
		v, err := fn(acc, el)
		if err != nil {
			return acc, err
		}
		acc = v
	}

	return acc, nil
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
