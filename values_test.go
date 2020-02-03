// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"regexp"
	"testing"
)

func TestInt(t *testing.T) {
	t.Parallel()

	f := Array(1, Int(1, 1))
	checkStringIs(t, f(nil).(Slice).String(), `[1]`, "single value Int")

	cases := []struct {
		gen    Generator
		expect []int
	}{
		{Array(1, Int(1, 3)), []int{1, 2, 3}},      // positive boundaries
		{Array(1, Int(-3, -1)), []int{-3, -2, -1}}, // negative boundaries
		{Array(1, Int(-1, 1)), []int{-1, 0, 1}},    // zero-spannning boundaries
	}

	for _, c := range cases {
		c := c
		checkFuncCoversIntRange(func() int { return c.gen(nil).(Slice)[0].(int) }, c.expect)
	}
}

func TestInt31(t *testing.T) {
	t.Parallel()

	f := Array(1, Int31(1, 1))
	checkStringIs(t, f(nil).(Slice).String(), `[1]`, "single value Int31")

	cases := []struct {
		gen    Generator
		expect []int32
	}{
		{Array(1, Int31(1, 3)), []int32{1, 2, 3}},      // positive boundaries
		{Array(1, Int31(-3, -1)), []int32{-3, -2, -1}}, // negative boundaries
		{Array(1, Int31(-1, 1)), []int32{-1, 0, 1}},    // zero-spannning boundaries
	}

	for _, c := range cases {
		c := c
		checkFuncCoversInt31Range(func() int32 { return c.gen(nil).(Slice)[0].(int32) }, c.expect)
	}
}

func TestFloat64(t *testing.T) {
	t.Parallel()

	f := Array(1, Float64(0, 0))
	checkStringIs(t, f(nil).(Slice).String(), `[0]`, "single value Float64")

	cases := []struct {
		low  float64
		high float64
	}{
		{0.0, 1.0},
		{0.0, 2.0},
		{-1.0, 1.0},
	}

	for _, c := range cases {
		f := Array(1, Float64(c.low, c.high))
		for i := 0; i < 10; i++ {
			n := f(nil).(Slice)[0].(float64)
			checkFloatInRange(t, n, c.low, c.high)
		}
	}
}

func TestDigits(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input string
		match string
	}{
		{``, `^$`},
		{`#`, `^\d$`},
		{`\#`, `^#$`},
		{`\\#`, `^\\\d$`},
		{`#-## #/#`, `^\d-\d\d \d/\d$`},
		{`###»###`, `^\d{3}»\d{3}$`},
	}

	for _, c := range cases {
		f := Array(1, Digits(c.input))
		s := f(nil).(Slice)[0].(string)
		match, err := regexp.MatchString(c.match, s)
		if err != nil {
			t.Errorf("match error: %v", err)
		} else if !match {
			t.Errorf("failed: `%s` doesn't match `%s`", s, c.match)
		}
	}
}

func TestHexDigits(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input string
		match string
	}{
		{``, `^$`},
		{`#`, `^[0-9a-f]$`},
		{`\#`, `^#$`},
		{`\\#`, `^\\[0-9a-f]$`},
		{`#-## #/#`, `^[0-9a-f]-[0-9a-f][0-9a-f] [0-9a-f]/[0-9a-f]$`},
		{`###»###`, `^[0-9a-f]{3}»[0-9a-f]{3}$`},
		{`############`, `^[0-9a-f]{12}$`},
	}

	for _, c := range cases {
		f := Array(1, HexDigits(c.input))
		s := f(nil).(Slice)[0].(string)
		match, err := regexp.MatchString(c.match, s)
		if err != nil {
			t.Errorf("match error: %v", err)
		} else if !match {
			t.Errorf("failed: `%s` doesn't match `%s`", s, c.match)
		}
	}
}
