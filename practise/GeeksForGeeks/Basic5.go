package main

import "fmt"

func main() {

	//creating an array
	arr := [10]string {
		"this","is","tutorial","of","go","programming","language",
	}
	fmt.Println(arr)

	mySlice := arr[2:6]
	fmt.Println(mySlice)
	fmt.Println("The capacity of the slice is",cap(mySlice))
	fmt.Println("The length of the slice is ",len(mySlice))

	mySlice = append(mySlice,"hello world")
	fmt.Println(mySlice,arr)
	fmt.Println("The capacity of the slice is",cap(mySlice))
	fmt.Println("The length of the slice is ",len(mySlice))

	mySlice = append(mySlice,"raghav")
	fmt.Println("The capacity of the slice is",cap(mySlice))
	fmt.Println("The length of the slice is ",len(mySlice))

	fmt.Println(arr)
}