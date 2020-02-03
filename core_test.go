// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"strings"
	"testing"
)

// checks for string equality
func checkStringIs(t *testing.T, got, wanted, label string) {
	t.Helper()
	if got != wanted {
		t.Errorf("%s failed: got `%s`; wanted `%s`", label, got, wanted)
	}
}

// checks float in range [low,high)
func checkFloatInRange(t *testing.T, got float64, low float64, high float64) {
	t.Helper()
	if got < low || got >= high {
		t.Errorf("check failed: `%v` not in range [%v, %v)", got, low, high)
	}
}

// Loop forever and get killed by timeout if not covered
func checkFuncCoversIntRange(f func() int, xs []int) {
	seen := make([]bool, len(xs))
	for {
		n := f()
		for i, x := range xs {
			if n == x {
				seen[i] = true
				break
			}
		}
		if countTrue(seen) == len(xs) {
			break
		}
	}
}

// Loop forever and get killed by timeout if not covered
func checkFuncCoversInt31Range(f func() int32, xs []int32) {
	seen := make([]bool, len(xs))
	for {
		n := f()
		for i, x := range xs {
			if n == x {
				seen[i] = true
				break
			}
		}
		if countTrue(seen) == len(xs) {
			break
		}
	}
}

func countTrue(xs []bool) int {
	sum := 0
	for _, x := range xs {
		if x {
			sum++
		}
	}
	return sum
}

func TestMap_String(t *testing.T) {
	t.Parallel()

	// Stringification of nil map
	var nilMap Map
	checkStringIs(t, nilMap.String(), "{}", "nil map")

	// Stringification of empty map
	emptyMap := Map{}
	checkStringIs(t, emptyMap.String(), "{}", "empty map")

	// Stringification of non-empty map
	regularMap := Map{
		"x": 42,
	}
	checkStringIs(t, regularMap.String(), `{"x":42}`, "empty map")

	// Stringification of un-marshalable Map
	badMap := Map{
		"x": func() {},
	}
	if s := badMap.String(); !strings.Contains(s, "could not marshal object") {
		t.Errorf("Didn't get expected marshaling error: got %q", s)
	}
}

func TestSlice_String(t *testing.T) {
	t.Parallel()

	// Stringification of nil array
	var nilSlice Slice
	checkStringIs(t, nilSlice.String(), "[]", "nil array")

	// Stringification of empty array
	emptySlice := Slice{}
	checkStringIs(t, emptySlice.String(), "[]", "empty array")

	// Stringification of non-empty array
	regularSlice := Slice{
		"x", 42,
	}
	checkStringIs(t, regularSlice.String(), `["x",42]`, "empty array")

	// Stringification of un-marshalable Slice
	badSlice := Slice{func() {}}
	if s := badSlice.String(); !strings.Contains(s, "could not marshal array") {
		t.Errorf("Didn't get expected marshaling error: got %q", s)
	}
}
