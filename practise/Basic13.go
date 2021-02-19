package main

import "fmt"

type Employee struct {
	name string
	age int
}

func (e Employee) updateName() {
	e.name = "Raghav Maheshwari"
}

//This should receive the value sent as a pointer.
func (e *Employee) updateName2(){
	e.name = "Raghav Maheshwari"
}

/**
	The function expects the same type sent.
	If address is sent, it will receive as a pointer otherise normal one
	In case of methods
	We can send either the address or the object itself.
 */
func updateName3(e *Employee){
	e.name = "Tony stark"
}

func main(){

	//We dont pass this by reference
	employee := Employee{name: "Raghav",age: 21}
	fmt.Println(employee)
	employee.updateName2()
	fmt.Println(employee)

	//To really change the value we need to pass it by value.
	employee2 := &Employee{name: "Rahul",age: 32}
	fmt.Println(employee2)
	fmt.Println(*employee2)
	employee2.updateName2()
	fmt.Println(*employee2)


	employee3 := Employee{name: "Tanisha",age: 34}
	fmt.Println(employee3)
	updateName3(&employee3)
	fmt.Println(employee3)
}

