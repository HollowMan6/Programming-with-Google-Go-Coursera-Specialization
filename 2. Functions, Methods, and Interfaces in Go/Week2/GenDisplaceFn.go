package main

import "fmt"

func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(a,v0,s0 float64) func(float64) float64{
	return func(t float64) float64 {
		return 1/2*a*t*t+v0*t+s0
	}
}