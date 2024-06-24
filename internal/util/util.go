package util

func Map[T, E any](arr []T, f func(T) E) []E {
	result := make([]E, 0, cap(arr))

	for _, elem := range arr {
		result = append(result, f(elem))
	}

	return result
}
