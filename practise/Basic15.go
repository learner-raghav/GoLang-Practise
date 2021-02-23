package main

import "fmt"
import "unsafe"
/**

	1. Arrays in Go cannot shrink to grow.
	2. They are completely static with predefined length.
	3. There are two ways to declare arrays in Golang

		1. var a [5]int
		2. b := [5]int{1,2,3,4,5}
 */


func appendCities(cities []string,city string){
	fmt.Println("Cities inside function!!")
	cities = append(cities,city)
	for i := range cities {
		fmt.Println(cities[i])
	}
	cities[0] = "Honolulu"
}

func main(){

	var a [10]int

	fmt.Println(a)
	fmt.Println(len(a))

	//As the statement below already gives the size as 80, this verifies that arrays are
	// allocated memory at compile time.
	fmt.Println(unsafe.Sizeof(a))
	for i := 0; i< 10;i ++ {
		a[i] = i + 1
	}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(unsafe.Sizeof(a))

	/**

		1. Slices are fancier arrays in Go, which means there are special commands,
			to, add additional values which make it grow, remove specific values to
			make it shrink and much more.

		2. Slices are initilaized in the same way as arrays are initialized, we just dont provide
		   the size and they grow.

		3. Slices are passed by reference by default!!
	 */

	cities := []string{"Ambala","Bangalore","Chennai","Kolkata"}
	for i := range cities {
		fmt.Println(cities[i])
	}
	fmt.Printf("%T",cities)

	fmt.Println("---- Append Operation---")
	cities = append(cities,"Delhi")
	for i := range cities {
		fmt.Println(cities[i])
	}

	appendCities(cities,"Vishakhapatnam")
	fmt.Println("----Printing outside the function!!")
	for i := range cities {
		fmt.Println(cities[i])
	}
}