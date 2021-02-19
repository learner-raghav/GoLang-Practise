package main

import "fmt"

/**

	A structure or struct in golang is a user defined type that allows to
	combine items of possibly different types into a single type.

	Any real world entity that has some properties or fields can be represented as a struct
	These can be called as lightweight classes that does not support inheritance but supports composition
 */

//Some application might need address of the users
type Address struct {
	name string
	street string
	city string
	state string
	pincode int
}

//We ave a compact way of defining structs too.
// If many fields have the same datatype we can simply group them

type Address1 struct {
	name, street, city, state string
	pincode int
}

func main(){

	// We declare an address type variable.
	var a Address
	// This gives default values for all the variables.
	fmt.Println(a)

	raghav := Address{"57","Ankit Vihar","Muzaffarnagar","Uttar Pradesh",251001}
	fmt.Println(raghav)

	//We can even name fields and define
	harshit := Address{name: "g-11",street: "MSR Nagar",pincode: 560054}
	fmt.Println(harshit)

	//Pointers to a struct
	employee := &Address{"g-12","MSR Nagar","Bangalore","Karnataka",560054}
	fmt.Println((*employee).pincode)
}