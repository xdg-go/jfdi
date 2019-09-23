// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package jfdi is short for JSON Fake Data Inventor.  It uses a declarative,
// functional syntax to produce randomly-generated, heterogeneous data
// structures.
//
// JSON, unlike Go, allows for data structures with mixed types.  These are
// most-easily modeled in Go with map[string]interface{} (for objects) and
// []interface{} (for arrays).  jfdi calls these Map and Slice, for short.
//
// jfdi uses higher-order functions extensively -- functions that return
// Generator functions, which are called to recursively produce the
// desired result.  For example, to define a JSON object, use the Object
// function with a template Map.  If any values in the Map are also
// Generators, they are recursively called to fill in the value:
//
//     factory := jfdi.Object( jfdi.Map{
//         "name"     : jfdi.Pick("Alice", "Bob", "Carol"),
//         "age"      : jfdi.Int(18, 65),
//         "ssn"      : jfdi.Digits("###-##-####"),
//         "friends"  : jfdi.Array(jfdi.Int(1,3), jfdi.Pick("Dan", "Eve", "Frank")),
//     })
//
//     object := factory(jfdi.NewContext())
//
//     fmt.Println(object)
//     // {"name":"Carol","age":42,"ssn":"314-15-9265","friends":["Eve"]}
//
// Define custom Generators or Generator constructors as needed if built-in
// Generator constructors aren't enough.
package jfdi
