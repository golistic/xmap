// Copyright (c) 2021, 2023, Geert JM Vanderkelen

package xmap

import (
	"sync"
)

// OrderedMap wraps around a Go map keeping the order with which
// elements have been added. Keys must be comparable, but values
// can be anything.
// Unlike map, index assigment is not possible. Use the `Set`
// method to set a key with a particular value.
// Use the Keys method to retrieves keys, Values to get the
// values. To get both, which probably what you want, use
// the KeysValues method.
type OrderedMap[T comparable] struct {
	mapMU sync.RWMutex
	map_  map[T]any
	order []T
}

// Count returns the number of elements in the map.
func (om *OrderedMap[T]) Count() int {
	return len(om.order)
}

// Set key in OrderedMap to value. Previously stored values
// are overwritten, but the order does not change.
func (om *OrderedMap[T]) Set(key T, value any) {
	om.mapMU.Lock()
	defer om.mapMU.Unlock()

	if om.map_ == nil {
		om.map_ = map[T]any{}
	}

	om.map_[key] = value
	if !om.has(key) {
		om.order = append(om.order, key)
	}
}

// Delete deletes the element with the specified key from
// the OrderedMap.
func (om *OrderedMap[T]) Delete(key T) {
	om.mapMU.Lock()
	defer om.mapMU.Unlock()

	for i := 0; i < len(om.order); i++ {
		if om.order[i] == key {
			om.order[i] = om.order[len(om.order)-1]
			om.order = om.order[:len(om.order)-1]
			delete(om.map_, key)
			break // key is unique; stop looking for more
		}
	}
}

// Keys returns keys as slice of string.
func (om *OrderedMap[T]) Keys() []T {
	om.mapMU.RLock()
	defer om.mapMU.RUnlock()

	return om.order
}

func (om *OrderedMap[T]) values() []any {
	res := make([]any, len(om.order))
	for i, k := range om.order {
		res[i] = om.map_[k]
	}
	return res
}

// Values returns the values as slice of interfaces.
func (om *OrderedMap[T]) Values() []any {
	om.mapMU.RLock()
	defer om.mapMU.RUnlock()

	return om.values()
}

// KeysValues returns the keys as slice of strings, and values as slice of interfaces.
func (om *OrderedMap[T]) KeysValues() ([]T, []any) {
	om.mapMU.RLock()
	defer om.mapMU.RUnlock()

	return om.order, om.values()
}

// Has returns whether the map contains key.
func (om *OrderedMap[T]) Has(key T) bool {
	om.mapMU.RLock()
	defer om.mapMU.RUnlock()

	return om.has(key)
}

func (om *OrderedMap[T]) has(key T) bool {
	for _, e := range om.order {
		if e == key {
			return true
		}
	}
	return false
}

// Value returns the value for key and also whether it was found.
// The bool is returned because value could be nil.
func (om *OrderedMap[T]) Value(key T) (any, bool) {
	om.mapMU.RLock()
	defer om.mapMU.RUnlock()

	return om.map_[key], om.has(key)
}
