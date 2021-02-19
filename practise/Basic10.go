package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64,float64) float64) float64 {
	return fn(3,4)
}
/**

	Function are values too. They can be passed around just like other values.
	The function values amy be used as function arguments and return values
 */
func main(){

	//This is how we create function values
	//Now, when we want, we can pass the values into hypotenuse too
	hypotenuse := func(x,y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}

	fmt.Println(hypotenuse(3,4))
	fmt.Println(compute(hypotenuse))
}
