package main

import (
	"fmt"
)

func main() {
	i := 0
	// Run forever to make it to show race condition
	for {
		var x, y int

		// Set x to 60
		go func(v *int) {
			*v = 60
		}(&x)

		// Set y to 3
		go func(v *int) {
			*v = 3
		}(&y)

		/*
		  Raca condition is when multiple threads are trying to access and manipulat the same variable.
		  the code below are all accessing and changing the value.
		  Divide x (60) by y (3) and assign to z (42)...
		  Except if y is not assigned 3 before x is assigned 60,
		  y's initialized value of 0 is used,
		  Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable,
		  which causes divide by zero exception.
		*/
		go func(v1 int, v2 int) {
			fmt.Println(v1 / v2)
		}(x, y)

		i += 1

		fmt.Printf("%d\n", i)
	}
}