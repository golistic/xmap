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
		got := xmap.Keys(m)
		sort.Strings(got)
		xt.Eq(t, exp, got)
	})

	t.Run("int keys", func(t *testing.T) {
		m := map[int]string{
			987: "bar",
			123: "foo",
		}

		exp := []int{123, 987}
		got := xmap.Keys(m)
		sort.Ints(got)

		xt.Eq(t, exp, got)
	})
}
