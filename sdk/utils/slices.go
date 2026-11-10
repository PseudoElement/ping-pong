package utils

func MergeSlices[T any](arrays [][]T) []T {
	var totalLen int
	for _, arr := range arrays {
		totalLen += len(arr)
	}

	finalSlice := make([]T, totalLen)
	var idx int
	for _, arr := range arrays {
		for _, el := range arr {
			finalSlice[idx] = el
			idx++
		}
	}

	return finalSlice
}
