// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Generator func(*Context) interface{}

var zeroGenerator = func(*Context) interface{} { return nil }

type Map map[string]interface{}

func (m Map) String() string {
	if m == nil {
		return "{}"
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("could not marshall object: %v", err)
	}
	return string(buf)
}

type Slice []interface{}

func (s Slice) String() string {
	if s == nil {
		return "[]"
	}
	buf, err := json.Marshal(s)
	if err != nil {
		return fmt.Sprintf("could not marshall array: %v", err)
	}
	return string(buf)
}

type Context struct {
	Depth int
	Rand  *rand.Rand
	Value Map
}

func NewContext() *Context {
	return &Context{
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
		Value: make(Map),
	}
}

func toGenerator(in interface{}) (Generator, bool) {
	if f, ok := in.(Generator); ok {
		return f, true
	} else if f, ok := in.(func(c *Context) interface{}); ok {
		return f, true
	} else if f, ok := in.(func() interface{}); ok {
		return func(*Context) interface{} { return f() }, true
	} else {
		return zeroGenerator, false
	}
}

func toMap(c *Context, in interface{}) (Map, bool) {
	if x, ok := in.(Map); ok {
		return x, true
	}
	if f, ok := toGenerator(in); ok {
		v := f(c)
		if x, ok := v.(Map); ok {
			return x, true
		}
	}
	return Map{}, false
}

func toInt(c *Context, in interface{}) (int, bool) {
	if x, ok := in.(int); ok {
		return x, true
	}
	if f, ok := toGenerator(in); ok {
		v := f(c)
		if x, ok := v.(int); ok {
			return x, true
		}
	}
	return 0, false
}

func toStr(c *Context, in interface{}) (string, bool) {
	if x, ok := in.(string); ok {
		return x, true
	}
	if f, ok := toGenerator(in); ok {
		v := f(c)
		if x, ok := v.(string); ok {
			return x, true
		}
	}
	return "", false
}

func expand(c *Context, in interface{}) interface{} {
	if f, ok := toGenerator(in); ok {
		return f(c)
	}
	return in
}

func mergeMaps(dst, src Map) {
	for k, v := range src {
		dst[k] = v
	}
}
