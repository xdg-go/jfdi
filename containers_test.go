// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestObject(t *testing.T) {
	t.Parallel()

	f := Object(Map{"x": 23, "y": func() interface{} { return 42 }})
	checkStringIs(t, f(nil).(Map).String(), `{"x":23,"y":42}`, "simple FakeHash")

	f = Object(Map{
		"x": 23,
		"y": func() interface{} { return 42 },
		"z": Object(Map{"a": 1, "b": "c"}),
	})
	checkStringIs(t, f(nil).(Map).String(),
		`{"x":23,"y":42,"z":{"a":1,"b":"c"}}`, "nested FakeHash")
}

func TestArray(t *testing.T) {
	t.Parallel()

	f := Array(2, 23)
	checkStringIs(t, f(nil).(Slice).String(), `[23,23]`, "simple FakeArray")

	f = Array(2, Array(3, 0))
	checkStringIs(t, f(nil).(Slice).String(), `[[0,0,0],[0,0,0]]`, "nested FakeArray")
}

func TestMaxDepth(t *testing.T) {
	t.Parallel()

	// Hash only
	f := MaxDepthObject(1, Map{"x": 23, "y": MaxDepthObject(1, Map{"z": 42})})
	checkStringIs(t, f(nil).(Map).String(), `{"x":23,"y":null}`, "simple FakeHash")

	f = MaxDepthObject(1, Map{"x": 23, "y": MaxDepthObject(2, Map{"z": 42})})
	checkStringIs(t, f(nil).(Map).String(), `{"x":23,"y":{"z":42}}`, "simple FakeHash")

	// Array only (in a hash wrapper)
	f = Object(Map{"x": MaxDepthArray(1, 3, "a")})
	checkStringIs(t, f(nil).(Map).String(), `{"x":null}`, "FakeArrayMaxDepth 1")

	f = Object(Map{"x": MaxDepthArray(2, 3, "a")})
	checkStringIs(t, f(nil).(Map).String(), `{"x":["a","a","a"]}`, "FakeArrayMaxDepth 2")

	// Mixed hash and array
	f = Object(Map{
		"x": MaxDepthArray(3, 1,
			MaxDepthObject(3, Map{
				"x": 42,
				"y": MaxDepthArray(3, 1, 42),
				"z": MaxDepthObject(3, Map{"x": 42}),
			}),
		),
	})
	checkStringIs(t, f(nil).(Map).String(),
		`{"x":[{"x":42,"y":null,"z":null}]}`, "hash+array, with depth limit")

	f = Object(Map{
		"x": MaxDepthArray(3, 1,
			MaxDepthObject(3, Map{
				"x": 42,
				"y": MaxDepthArray(9, 1, 42),
				"z": MaxDepthObject(9, Map{"x": 42}),
			}),
		),
	})
	checkStringIs(t, f(nil).(Map).String(),
		`{"x":[{"x":42,"y":[42],"z":{"x":42}}]}`, "hash+array, without depth limit")
}

func TestSequence(t *testing.T) {
	t.Parallel()

	f := Sequence(2, 23)
	checkStringIs(t, f(nil).(Slice).String(), `[2,23]`, "simple FakeSequence")

	f = Sequence(2, Sequence(3, 4))
	checkStringIs(t, f(nil).(Slice).String(), `[2,[3,4]]`, "nested FakeSequence")
}

func TestSeededMapMerge(t *testing.T) {
	t.Parallel()

	f := Object(Map{"x": Int(1, 1000)}, Map{"y": Int(1, 1000)})

	getSeededContext := func() *Context {
		c := NewContext()
		c.Rand = rand.New(rand.NewSource(42))
		return c
	}

	// Expect the same results for two series of maps generated from the same
	// seed.

	var first []string
	var second []string

	c1 := getSeededContext()
	c2 := getSeededContext()

	for i := 0; i < 10; i++ {
		first = append(first, f(c1).(Map).String())
		second = append(second, f(c2).(Map).String())
	}

	for i := 0; i < 10; i++ {
		checkStringIs(t, first[i], second[i], fmt.Sprintf("seeded map merge, doc %d", i))
	}
}
