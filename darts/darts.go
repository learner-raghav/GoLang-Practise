package darts

import "math"

func Score(x float64, y float64) int{
	ans := math.Sqrt(x*x+y*y)
	if ans <= 1.0 {
		return 10
	}
	if ans <= 5.0 {
		return 5
	}
	if ans <= 10.0 {
		return 1
	}
	return 0
}