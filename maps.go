package go_utils

import (
	"encoding/json"
	"fmt"
)

// CompareMapsKeys returns true if the keys of two maps are equal.
func CompareMapsKeys[kt comparable, vt comparable](a, b map[kt]vt) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}

	return true
}

// CompareMapsValues returns true if the values of two maps are equal.
func CompareMapsValues[kt comparable, vt comparable](a, b map[kt]vt) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}

	return true
}

// CompareMaps returns true if two maps are equal.
func CompareMaps[kt comparable, vt comparable](a, b map[kt]vt) bool {
	return CompareMapsKeys(a, b) && CompareMapsValues(a, b)
}

// GetMapKeys returns the keys of a map.
func GetMapKeys[kt comparable, vt interface{}](m map[kt]vt) []kt {
	keys := make([]kt, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// GetMapValues returns the values of a map.
func GetMapValues[kt comparable, vt comparable](m map[kt]vt) []vt {
	values := make([]vt, len(m))

	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

// CopyMap copies a map.
func CopyMap[kt comparable, vt interface{}](m map[kt]vt) map[kt]vt {
	mCpy := make(map[kt]vt)

	for k, v := range m {
		mCpy[k] = v
	}

	return mCpy
}

func PrintMap(mp map[any]any) {
	b, err := json.MarshalIndent(mp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
}
