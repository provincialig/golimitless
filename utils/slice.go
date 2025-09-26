package utils

func SliceFilter[T any](arr []T, fn func(el T) bool) []T {
	result := []T{}

	for _, el := range arr {
		if fn(el) {
			result = append(result, el)
		}
	}

	return result
}

func SliceMap[T any, K any](arr []T, fn func(el T) (K, error)) ([]K, error) {
	result := []K{}

	for _, el := range arr {
		r, err := fn(el)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}
