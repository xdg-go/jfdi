package examples

import (
	"encoding/json"
	"fmt"

	"github.com/xdg-go/jfdi"
)

func Example_correlatedKeys() {
	// CorrelatedKeys shows how to write a custom generator that enforces a
	// relationship between randomly generated data.  The idea here is to use
	// the passed-in context to generate independent data values from other
	// generators and then construct a data structure with the right
	// relationship.

	// minMaxPair is a generator that creates a map with two keys,
	// "firstStudent" and "lastStudent", where the names are correctly ordered
	// despite being randomly generated.  (Duplicates are allowed in this simple
	// example.)
	minMaxPair := func(ctx *jfdi.Context) interface{} {
		studentNames := jfdi.Pick("Alice", "Bob", "Carol")

		// Generate two names
		min := studentNames(ctx).(string)
		max := studentNames(ctx).(string)

		// Reorder them, if needed,
		if min > max {
			min, max = max, min
		}

		// Return them as a Map
		return jfdi.Map{
			"firstStudent": min,
			"lastStudent":  max,
		}
	}

	// This factory is an Object generator that merges a Map generating a
	// classroom key with the Map containing the student names.
	factory := jfdi.Object(
		jfdi.Map{
			"classroom": jfdi.Pick("101", "101", "103"),
		},
		minMaxPair,
	)

	object := factory(jfdi.NewContext())
	output, _ := json.Marshal(object)
	fmt.Println(string(output))

	// Example output:
	// {"classroom":"101","firstStudent":"Bob","lastStudent":"Carol"}
}
