package main

/**
	closures in golang

	1. Go lang has a special feature called an anonymous function. An anonymous function can form a closure.
	2. A closure is a special type of anonymous function that references variables declared outside of the function itself.
*/
import "fmt"
func main(){

	GFG := 0
	counter := func() int {
		GFG += 1
		return GFG
	}

	fmt.Println(counter())
}