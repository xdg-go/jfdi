// Copyright 2019 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package jfdi

import "testing"

func TestPick(t *testing.T) {
	t.Parallel()

	f := Array(1, Pick())
	checkStringIs(t, f(nil).(Slice).String(), `[null]`, "empty Pick")

	f = Array(1, Pick(23))
	checkStringIs(t, f(nil).(Slice).String(), `[23]`, "single value Pick")

	f = Pick(23, 42)
	checkFuncCoversIntRange(func() int { return f(nil).(int) }, []int{23, 42})

	f = Pick(Int(1, 1), Int(2, 2), Int(3, 3))
	checkFuncCoversIntRange(func() int { return f(nil).(int) }, []int{1, 2, 3})
}
