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

func Word() Generator {
	f := Words(1)
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return f(c).([]string)[0]
	}
}

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

func Sentence() Generator {
	f := Sentences(1)
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return f(c).([]string)[0]
	}
}

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
