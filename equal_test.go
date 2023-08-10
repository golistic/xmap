// Copyright (c) 2023, Geert JM Vanderkelen

package xmap_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/golistic/xt"

	"github.com/golistic/xmap"
)

func TestEqual(t *testing.T) {
	t.Run("lengths are not equal", func(t *testing.T) {
		m1 := map[string]int{"1": 1}
		m2 := map[string]int{"1": 1, "2": 2}
		xt.Assert(t, !xmap.Equal(m1, m2))
	})

	t.Run("lengths are not equal", func(t *testing.T) {
		m1 := map[string]string{"1": "1"}
		m2 := map[string]string{"1": "1"}
		xt.Assert(t, xmap.Equal(m1, m2))
	})

	t.Run("values are not equal", func(t *testing.T) {
		m1 := map[float64]string{1.2: "1"}
		m2 := map[float64]string{1.2: "1123"}
		xt.Assert(t, !xmap.Equal(m1, m2))
	})

	t.Run("values are equal", func(t *testing.T) {
		m1 := map[uint]string{2: "something", 1: "1123"}
		m2 := map[uint]string{1: "1123", 2: "something"}
		xt.Assert(t, xmap.Equal(m1, m2))
		xt.Assert(t, reflect.DeepEqual(m1, m2))
	})

	t.Run("cannot handle maps", func(t *testing.T) {
		type rt map[int]string

		r := rt{1: "1"}

		m1 := map[string]any{
			"1": r,
		}

		m2 := map[string]any{
			"1": r,
		}

		xt.Panics(t, func() {
			xmap.Equal(m1, m2)
		})
	})
}

func BenchmarkEqual(b *testing.B) {
	v := strings.Repeat("a", 10)
	m1, m2 := map[int]string{}, map[int]string{}

	for i, j := 0, 100; i < 100; i, j = i+1, j-1 {
		m1[i] = v
		m2[j] = v
	}

	b.Run("using our Equal", func(b *testing.B) {
		xmap.Equal(m1, m2)
	})

	b.Run("using our reflect.DeepEqual", func(b *testing.B) {
		reflect.DeepEqual(m1, m2)
	})
}
