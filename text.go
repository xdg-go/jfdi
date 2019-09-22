// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"strings"

	"github.com/icrowley/fake"
)

// Word returns a generator that produces a randomly-chosen latin word.
func Word() Generator {
	f := Words(1)
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return f(c).([]string)[0]
	}
}

// Words returns a generator that produces a slice of randomly-chosen latin word
// of a given length.  The argument must be an integer or a generator of
// integers.  If the length is negative or zero, the generator panics.
func Words(n interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		length, ok := toInt(c, n)
		if !ok || length < 0 {
			panic("length must be a non-negative int or generate a non-negative int")
		}
		output := make([]string, length)
		for i := 0; i < length; i++ {
			output[i] = fake.Word()
		}
		return output
	}
}

// Sentence returns a generator that produces a randomly-generated 'latin sentence'.
func Sentence() Generator {
	f := Sentences(1)
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return f(c).([]string)[0]
	}
}

// Sentences returns a generator that produces a slice of randomly-generated
// latin 'sentences' of a given slice length.  The argument must be an integer
// or a generator of integers.  If the length is negative or zero, the
// generator panics.
func Sentences(n interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		length, ok := toInt(c, n)
		if !ok || length < 0 {
			panic("length must be a non-negative int or generate a non-negative int")
		}
		output := make([]string, length)
		for i := 0; i < length; i++ {
			output[i] = strings.Title(fake.Word()) + fake.Sentence()
		}
		return output
	}
}

// Join returns a generator that joins a slice of strings with a separator
// string.  The first argument must be a slice of string or a generator of
// them; the second argument must be a string or a generator of one.   The
// Generator panics if either argument produces an invalid type.
func Join(inputs interface{}, separator interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		sep, ok := toStr(c, separator)
		if !ok {
			panic("separator must be or generate a string")
		}
		xs := expand(c, inputs)
		ys, ok := xs.([]string)
		if !ok {
			panic("inputs must be or generate a slice of string")
		}
		return strings.Join(ys, sep)
	}
}
