// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import (
	"strings"
	"unicode/utf8"
)

// Int returns a generator that a random integer in the range [low,high].  If
// `low` is greater than `high`, it panics.
func Int(low, high int) Generator {
	if low > high {
		panic("first argument must be <= second argument")
	}
	if low == high {
		return func(c *Context) interface{} {
			return low
		}
	}
	span := high - low + 1
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return low + c.Rand.Intn(span)
	}
}

// Float64 returns a generator that a random float64 in the range [low,high).
// If `low` is greater than `high`, it panics.
func Float64(low, high float64) Generator {
	if low > high {
		panic("first argument must be <= second argument")
	}
	if low == high {
		return func(c *Context) interface{} {
			return low
		}
	}
	span := high - low
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		return low + span*c.Rand.Float64()
	}
}

var hexDigits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// Digits returns a generator that replaces `#` characters in a template string
// with a random digit from 0 to 9.  Backslashes will be treated as escape
// characters.
func Digits(pattern string) Generator {
	return RuneMap(pattern, func(c *Context, r rune) rune {
		if r == '#' {
			return hexDigits[c.Rand.Intn(10)]
		}
		return r
	})
}

// HexDigits returns a generator that replaces `#` characters in a template
// string with a random hexadecimal digit from 0 to f.  Backslashes will be
// treated as escape characters.
func HexDigits(pattern string) Generator {
	return RuneMap(pattern, func(c *Context, r rune) rune {
		if r == '#' {
			return hexDigits[c.Rand.Intn(16)]
		}
		return r
	})
}

// RuneMap returns a generator that replaces runes in a template pattern via a
// user-defined replacement function.  Runes in the pattern may be
// backslash-escaped to prevent replacement.
func RuneMap(pattern string, replacer func(*Context, rune) rune) Generator {
	return func(c *Context) interface{} {
		if c == nil {
			c = NewContext()
		}
		if len(pattern) == 0 {
			return ""
		}
		output := strings.Builder{}
		var r rune
		for i, w := 0, 0; i < len(pattern); i += w {
			r, w = utf8.DecodeRuneInString(pattern[i:])
			switch r {
			case '\\':
				// skip this and write next character
				r2, w2 := utf8.DecodeRuneInString(pattern[i+w:])
				w += w2
				output.WriteRune(r2)
			default:
				output.WriteRune(replacer(c, r))
			}
		}
		return output.String()
	}
}
