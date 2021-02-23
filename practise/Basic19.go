package main

import "fmt"

type student1 struct {
	firstName string
	lastName string
}

func main(){
	stu := student1{"raghav","maheshwari"}
	fmt.Println(stu)
	studentFunc1(stu)
	fmt.Println(stu)
	stu.studentFunc2()
	fmt.Println(stu)

	studentFunc3(&stu)
	fmt.Println(stu)

	stu.studentFunc4()
	fmt.Println(stu)
}

func (stu *student1) studentFunc4() {
	fmt.Println("----Enters function 4----")
	fmt.Println(stu)
	stu.lastName = "C Patel"
	fmt.Println(stu)
	fmt.Println("---Exits function 4---")
}

/**
	The output here clearly indicates that that structs are passed by value
	ans they are not passed by reference in function arguments as well as receivers.

	The outout also indicates that if we pass by refernce, we can change values in functions as well as a receiver.
	We need to receive as a reference in the receiver if we aim to change trhe values.

	So, basically structs by default are passed by VALUE and not by REFERENCE.

	In Go, if we want to take the address of the variable we are using
	, then we do not need to specify at the calling end. Specifying the nature of
	the passing(value or reference) at the function is sufficient.
*/

func studentFunc3(stu *student1) {
	fmt.Println("----Enters function 3----")
	fmt.Println(stu)
	stu.lastName = "Patel"
	fmt.Println(stu)
	fmt.Println("---Exits function 3---")
}

func (stu student1) studentFunc2(){
	fmt.Println("----Enters function 2----")
	fmt.Println(stu)
	stu.lastName = "Patel"
	fmt.Println(stu)
	fmt.Println("---Exits function 2---")

}

func studentFunc1(stu student1){
	fmt.Println("----Enters function----")
	fmt.Println(stu)
	stu.lastName = "Patel"
	fmt.Println(stu)
	fmt.Println("---Exits function---")
}

