package main

import (
	"fmt"
	"math"
)

/**

	1. Go does not have classes, However we can define methods on type.
	2. A method is a function with a special receiver argument.
	3. The receiver appears in its own argument list between the func keyword and the method name
 */

type Vertex struct{
	X,Y float64
}

//Written with the use of receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Y*v.Y + v.X*v.X)
}

// Written using a normal function
func Abs1(v Vertex) float64 {
	return math.Sqrt(v.Y*v.Y+v.X*v.X)
}

//Now, it tells that we can declare methods on non struct types too.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else{
		return float64(f)
	}
}

/**
	This is one of the most collest piece of code iwill come across
 */

func (v *Vertex) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func main(){
	v := Vertex{3,4}
	//fmt.Println(v.Abs())
	//fmt.Println(Abs1(v))
	v.Scale(10.0)
	fmt.Println(v)

}

