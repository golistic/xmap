// Copyright (c) 2023, Geert JM Vanderkelen

package xmap

func Keys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
