// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import "sort"

// Object consumes a variable number of Maps or Map generators to produce a new
// Generator that constructs a Map.  The input list of (possibly generated)
// Maps are merged by key (last key wins), and then any Generators among the
// Map values are replaced by the output of the Generator.
//
// If any argument is not a Map or Map generator, the function will panic.
func Object(xs ...interface{}) Generator {
	return MaxDepthObject(0, xs...)
}

// MaxDepthObject works like Object, but it takes an initial argument indicating
// a maximum depth in a compound data structure.  The resulting Generator will
// return nil instead of a Map if the maxDepth is exceeded.  A maxDepth of 0
// means depth is unlimited.
func MaxDepthObject(maxDepth int, xs ...interface{}) Generator {
	// Generator constructs an empty Map by iterating over keys of input map.
	// Each key corresponds to either a value or a Generator.  If its a
	// Generator, get the output value from it.
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		c.Depth++
		if maxDepth > 0 && c.Depth > maxDepth {
			return nil
		}

		// Build up model from inputs. Inputs must be maps or Generators of maps.
		model := Map{}
		for _, x := range xs {
			if m, ok := toMap(c, x); ok {
				mergeMaps(model, m)
			} else {
				panic("arguments must be Maps or generators of Maps")
			}
		}

		// Call expand by key in sorted order to ensure determinism
		// in random number generation associated with each key.
		output := Map{}
		keys := make([]string, 0, len(model))
		for k := range model {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			output[k] = expand(c, model[k])
		}
		return output
	}
}

// Array returns a Generator that constructs a Slice.  The first argument, which
// must be an int or an int generator, determines the length of the Slice; the
// second argument, which may be a value or an arbitrary Generator, is used for
// elements of the Slice.
//
// By mixing Generators or constants, various Slice structures are possible:
//   jfdi.Array(3,             42)            // [42, 42, 42]
//   jfdi.Array(3,             jfdi.Int(1,6)) // 3 elements of integers from 1-6
//   jfdi.Array(jfdi.Int(2,4), jfdi.Int(1,6)) // 2-4 elements of integers from 1-6
func Array(length, elementModel interface{}) Generator {
	return MaxDepthArray(0, length, elementModel)
}

// MaxDepthArray works like Array, but it takes an initial argument indicating
// a maximum depth in a compound data structure.  The resulting Generator will
// return nil instead of an Array if the maxDepth is exceeded.  A maxDepth of 0
// means depth is unlimited.
func MaxDepthArray(maxDepth int, length, elementModel interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		c.Depth++
		if maxDepth > 0 && c.Depth > maxDepth {
			return nil
		}

		n, ok := toInt(c, length)
		if !ok || n < 0 {
			panic("length must be a non-negative int or generate a non-negative int")
		}
		output := make(Slice, n)
		for i := 0; i < n; i++ {
			output[i] = expand(c, elementModel)
		}
		return output
	}
}

// Sequence returns a Generator that constructs a Slice. Each elementModel, which may be a value
// or a Generator, will be used as elements of the Slice at their respective positions.
//
// By mixing Generators or constants, various Slice structures are possible:
//   jfdi.Sequence(3, 42)            // [3, 42]
//   jfdi.Sequence(jfdi.Int(1,3), jfdi.Int(4,6)) // 2 elements of integers between 1-3 and 4-6 respectively.
func Sequence(elementModels ...interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}

		c.Depth++
		output := make(Slice, len(elementModels))
		for i := 0; i < len(elementModels); i++ {
			output[i] = expand(c, elementModels[i])
		}
		return output
	}
}
