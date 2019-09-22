// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

// Pick returns a generator that chooses one of the arguments with uniform
// liklihood.  If the chosen item is a Generator, the value produced by that
// generator is returned instead.  If no arguments are provided, the generator
// returns nil.
func Pick(xs ...interface{}) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		if len(xs) > 0 {
			return expand(c, xs[c.Rand.Intn(len(xs))])
		}
		return nil
	}
}
