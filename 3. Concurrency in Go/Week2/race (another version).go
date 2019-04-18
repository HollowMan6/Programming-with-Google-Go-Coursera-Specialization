package main

import (
  "fmt"
)

/*
Raca condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, add_one and sub_one are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.

*/


func add_one(pt *int) {
  (*pt)++
  fmt.Println(*pt)
}

func sub_one(pt *int) {
  (*pt)--
  fmt.Println(*pt)
}
func main() {
  i := 0
  go add_one(&i)
  go sub_one(&i)

  i++
  fmt.Println()

}
