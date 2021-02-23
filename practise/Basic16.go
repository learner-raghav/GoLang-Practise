package main

import "fmt"

/**

	1. Custom Type - We can define our own custom types which is cities here and we can
		use that in the code as it is easier to use.
	2.
 */

//This defines a type alias which is nothing but a string slice.
type cities []string

func main(){

	india := cities{"Delhi","Bangalore","Chennai","Mumbai"}

	//we call the receiver function on the type alias we created.
	india.Print()
}

//Here it is a receiver function.
//These functions are attached to the variable of type cities.
func (city cities) Print(){
	for i := range city {
		fmt.Println(city[i])
	}
}