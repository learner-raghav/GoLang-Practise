package main

import "fmt"

/**
	1. Before learning anything about what interfaces are, lets take a look at an example,
		to clearly get the idea regarding why we need interfaces and how it could save
		us a lot of trouble while coding in GO!!

	2. As we can see we had to write 2 functions with different receiver types. That
		is not a good coding practise. So, to get around this, we have the concept of interfaces.

	3. We couldn't have passed different arguments to functions as go is a strictly typed language.
 */

//We create an interface and defined the method there.
/**
	1. This interface basically tells GO, that all variables which have a receiver function
		which is named as printMarks, they all belong to type marks.

	2. So, interface is a type in GO that incorporates multiple types.
 */
type marks interface{
	printMarks()
}

type grade11Marks struct {
	math int
	physics int
}

type grade12Marks struct {
	math int
	computer int
}

func main(){
	ram := grade11Marks{56,67}
	shyam := grade12Marks{67,78}
	fmt.Printf("%T ",ram)
	fmt.Printf("%T ",shyam)
	print(ram)
	print(shyam)
}

//As grade11Marks implement the function, it is of type marks
func (ram grade11Marks) printMarks(){
	fmt.Println(ram.math," ",ram.physics)
}

//This will also be of type marks
func (shyam grade12Marks) printMarks(){
	fmt.Println(shyam.math," ",shyam.computer)
}

func print(b marks){
	b.printMarks()
}