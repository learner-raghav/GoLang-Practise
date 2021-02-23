package main

import "fmt"

type student struct{
	firstName, lastName string
	age int
}

func main(){

	/**
		Structs:

		1. Structs are a typed collection of fields. Basically, a group of information
			are grouped together as a type, which can be used to create instances of
			the struct we defined.

		2. Suppose, we need to have details of a pewrson, we cant simply have
			variables for the name, age and so on,

		3. A struct is definitely a better data structure for the same.
	 */

	raghav := student{"Raghav","Maheshwari",23}
	nimish := student{"Nimish","Bongale",34}

	fmt.Println(raghav)
	fmt.Println(nimish)

	raghav.Print()
	nimish.Print()
}

func (stu student ) Print(){
	fmt.Println(stu.age," ",stu.firstName," ",stu.lastName)
}
