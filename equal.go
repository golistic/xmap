// Copyright (c) 2023, Geert JM Vanderkelen

package xmap

// Equal returns whether maps m1 and m2 are equal. Returns false if
// length is not equal, key of m1 is not present in m2, or when value
// is not equal for matching keys.
// This is the simple case where we only allow comparable types as defined
// by Go.
// See this as a stricter, more specialized, but faster version of
// `reflect.DeepEqual` (when comparing maps).
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}
