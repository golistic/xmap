// Copyright (c) 2023, Geert JM Vanderkelen

package xmap_test

import (
	"sort"
	"testing"

	"github.com/golistic/xt"

	"github.com/golistic/xmap"
)

func TestKeys(t *testing.T) {
	t.Run("string keys", func(t *testing.T) {
		m := map[string]int{
			"foo": 123,
			"bar": 987,
		}

		exp := []string{"bar", "foo"}
		have := xmap.Keys(m)
		sort.Strings(have)
		xt.Eq(t, exp, have)
	})

	t.Run("int keys", func(t *testing.T) {
		m := map[int]string{
			987: "bar",
			123: "foo",
		}

		exp := []int{123, 987}
		have := xmap.Keys(m)
		sort.Ints(have)

		xt.Eq(t, exp, have)
	})
}
