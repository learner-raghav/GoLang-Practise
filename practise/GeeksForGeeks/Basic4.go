package main

import (
	"fmt"
	"strings"
)

/**
	Program to count the frequenct of each word in a string
 */

func main() {

	input_string := "hello from raghav and hello from raghav bangalore"
	input_string_array := strings.Split(input_string," ")
	fmt.Println(input_string_array)

	//defining the map to get the count
	var a =  make(map[string]int)
	//
	//a["raghav"] = 2
	//fmt.Println(a)
	//a["rahul"] = 3
	//fmt.Println(a)
	//
	//ans,ok := a["raghavs"]
	//fmt.Println(ans)
	//fmt.Println(ok)
	for i := range input_string_array {
		ans,ok  := a[input_string_array[i]]
		if ok {
			a[input_string_array[i]] = ans + 1
		} else {
			a[input_string_array[i]] = 1
		}
	}



	fmt.Println(a)
}
