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

func TestWords(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`^\S+$`)

	w := Word()(nil).(string)
	if !re.MatchString(w) {
		t.Errorf("failed: `%s` doesn't match `%s`", w, re)
	}

	for i := 0; i <= 10; i++ {
		f := Words(i)
		s := f(nil).([]string)
		if len(s) != i {
			t.Errorf("Words(%d) produces slice of length %d", i, len(s))
		}
		for _, v := range s {
			if !re.MatchString(v) {
				t.Errorf("failed: `%s` doesn't match `%s`", v, re)
			}
		}
	}

	xs := Words(Int(3, 5))(nil).([]string)
	if len(xs) < 3 || len(xs) > 5 {
		t.Errorf("Words(Int(x,y)) returned incorrect length slice")
	}
	for _, v := range xs {
		if !re.MatchString(v) {
			t.Errorf("failed: `%s` doesn't match `%s`", v, re)
		}
	}
}

func TestSentences(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`^[A-Z]\S*(\s*\S+)*[.!?]$`)

	s := Sentence()(nil).(string)
	if !re.MatchString(s) {
		t.Errorf("failed: `%s` doesn't match `%s`", s, re)
	}

	for i := 0; i <= 10; i++ {
		f := Sentences(i)
		s := f(nil).([]string)
		if len(s) != i {
			t.Errorf("Sentences(%d) produces slice of length %d", i, len(s))
		}
		for _, v := range s {
			if !re.MatchString(v) {
				t.Errorf("failed: `%s` doesn't match `%s`", v, re)
			}

		}
	}

	xs := Sentences(Int(3, 5))(nil).([]string)
	if len(xs) < 3 || len(xs) > 5 {
		t.Errorf("Sentences(Int(x,y)) returned incorrect length slice")
	}
	for _, v := range xs {
		if !re.MatchString(v) {
			t.Errorf("failed: `%s` doesn't match `%s`", v, re)
		}
	}
}

func TestJoin(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input interface{}
		sep   interface{}
		match string
	}{
		{[]string{"a", "b", "c"}, " ", `^a b c$`},
		{[]string{"a", "b", "c"}, Pick(" ", "-"), `^a[- ]b[- ]c$`},
		{Words(3), " ", `^\S+ \S+ \S+$`},
		{Words(3), "", `^\S+$`},
		{Words(3), Pick(" ", "-"), `^\S+[- ]\S+[- ]\S+$`},
		{Words(0), "", `^$`},
	}

	for _, c := range cases {
		f := Join(c.input, c.sep)
		s := f(nil).(string)
		re := regexp.MustCompile(c.match)
		if !re.MatchString(s) {
			t.Errorf("failed: `%s` doesn't match `%s`", s, c.match)
		}
	}
}
