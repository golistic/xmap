// Copyright (c) 2023, Geert JM Vanderkelen

package xmap

// Keys returns the keys of the map m.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {

	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
