# jfdi - JSON Fake Data Inventor (experimental)

[![GoDoc](https://godoc.org/github.com/xdg-go/jfdi?status.svg)](https://godoc.org/github.com/xdg-go/jfdi) [![Build Status](https://travis-ci.org/xdg-go/jfdi.svg?branch=master)](https://travis-ci.org/xdg-go/jfdi) [![codecov](https://codecov.io/gh/xdg-go/jfdi/branch/master/graph/badge.svg)](https://codecov.io/gh/xdg-go/jfdi) [![Go Report Card](https://goreportcard.com/badge/github.com/xdg-go/jfdi)](https://goreportcard.com/report/github.com/xdg-go/jfdi) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A Go library for declarative, randomly-generated, heterogeneous JSON data
structures.

[Still under development.]

JSON, unlike Go, allows for data structures with mixed types.  These are
most-easily modeled in Go with `map[string]interface{}` (for objects) and
`[]interface{}` (for arrays).  `jfdi` calls these `Map` and `Slice`, for short.

`jfdi` uses higher-order functions extensively -- functions that return
`Generator` functions, which are called to recursively produce the
desired result.  For example, to define a JSON object, use the `Object`
function with a template `Map`.  If any values in the `Map` are also
`Generator`s, they are recursively called to fill in the value:

    factory := jfdi.Object( jfdi.Map{
        "name"     : jfdi.Pick("Alice", "Bob", "Carol"),
        "age"      : jfdi.Int(18, 65),
        "ssn"      : jfdi.Digits("###-##-####"),
        "friends"  : jfdi.Array(jfdi.Int(1,3), jfdi.Pick("Dan", "Eve", "Frank")),
    })

    object := factory(jfdi.NewContext())

    fmt.Println(object)
    // {"name":"Carol","age":42,"ssn":"314-15-9265","friends":["Eve"]}

Define custom `Generator`s or `Generator` constructors as needed if built-in
`Generator` constructors aren't enough.

# Copyright and License

Copyright 2019 by David A. Golden. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
