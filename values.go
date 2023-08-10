// Copyright (c) 2023, Geert JM Vanderkelen

package xmap

// Values returns the values of the map m.
func Values[M ~map[K]V, K comparable, V any](m M) []V {

	keys := make([]V, 0, len(m))

	for _, v := range m {
		keys = append(keys, v)
	}

	return keys
}
