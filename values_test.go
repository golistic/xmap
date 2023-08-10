// Copyright (c) 2023, Geert JM Vanderkelen

package xmap_test

import (
	"sort"
	"testing"

	"github.com/golistic/xt"

	"github.com/golistic/xmap"
)

func TestValues(t *testing.T) {
	t.Run("int values", func(t *testing.T) {
		m := map[string]int{
			"foo":    123,
			"bar":    987,
			"foobar": 2,
		}

		exp := []int{2, 123, 987}
		got := xmap.Values(m)
		sort.Ints(got)
		xt.Eq(t, exp, got)
	})

	t.Run("string values", func(t *testing.T) {
		m := map[int]string{
			987: "bar",
			123: "foo",
			2:   "a",
		}

		exp := []string{"a", "bar", "foo"}
		got := xmap.Values(m)
		sort.Strings(got)

		xt.Eq(t, exp, got)
	})
}
