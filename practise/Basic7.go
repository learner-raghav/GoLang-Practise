package main

import "fmt"

func main() {

	//simply declaring these variables.
	var p *int
	i := 42
	p = &i //We are storing the address of variable i in p.

	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	// A pointer to P will now give the value of i

	j:= 20
	addressOfJ := &j

	fmt.Println(j)
	fmt.Println(addressOfJ)
	fmt.Println(*addressOfJ)

	*addressOfJ = 34
	fmt.Println(addressOfJ)
	fmt.Println(j)
	fmt.Println(*addressOfJ)
}