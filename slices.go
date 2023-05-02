package go_utils

// DeleteFromSliceByValue deletes the first occurrence of an element from a slice.
func DeleteFromSliceByValue[T comparable](slc []T, el T) {
	for i, v := range slc {
		if v == el {
			slc = append(slc[:i], slc[i+1:]...)
		}
	}
}

// DeleteFromSliceByIndex deletes an element from a slice by its index.
func DeleteFromSliceByIndex[T interface{}](slc []T, index int) {
	slc = append(slc[:index], slc[index+1:]...)
}

// CopySlice copies a slice.
func CopySlice[T interface{}](slc []T) []T {
	slcCpy := make([]T, len(slc))
	copy(slcCpy, slc)

	return slcCpy
}

// GetSlicesIntersection returns the intersection of two slices.
func GetSlicesIntersection[T comparable](a, b []T) []T {
	var frequencyMap = make(map[T]int)
	var result []T

	for _, v := range a {
		frequencyMap[v]++
	}

	for _, v := range b {
		if frequencyMap[v] > 0 {
			result = append(result, v)
		}
	}

	return result
}

// GetSlicesUnion returns the union of two slices.
func GetSlicesUnion[T comparable](a, b []T) []T {
	var frequencyMapForA = make(map[T]int)
	var frequencyMapForB = make(map[T]int)

	var result []T

	for _, v := range a {
		frequencyMapForA[v]++
	}

	for _, v := range b {
		frequencyMapForB[v]++
	}

	for _, v := range a {
		if frequencyMapForB[v] == 0 {
			result = append(result, v)
		}
	}

	for _, v := range b {
		if frequencyMapForA[v] == 0 {
			result = append(result, v)
		}
	}

	return result
}
