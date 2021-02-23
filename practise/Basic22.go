package main

import "fmt"

/**
	1. Go language interfaces are different from other languages. In Go language,
		the interface is a custom type that is used to specify a set of one or more method
		signatures and the interface is abstract, so we are not allowed to create an instance of an interface.

	2. But we can create a variable of type interface and this variable can be assigned with a concrete type value that has methods the interface requires.
		Or in other words, we can say that an interface is a collection of methods as well as it is a custom type.

	3. In Go language, it is necessary to implement all the methods declared in the interface for implementing
		an interface. The go language interfaces are implemented implicitly. And it does
		not contain any specific keyword to implement interfaces just like other languages.

	4. The zero value of the interface is nil.

	5. When an interface contains zero methods, such types of interfaces are known as empty
		interface and hence all types implement the empty interface.

	6. We can use interfaces when multiple struct types have the same method.
 */

type tank interface {
	//It will have an area and a volume
	area() float64
	volume() float64
}

type myValue struct {
	radius float64
	height float64
}

func (mVal myValue) area() float64{
	return 2*mVal.height*mVal.radius
}

func (mVal myValue) volume() float64 {
	return 3*mVal.radius*mVal.height
}

func main(){
	var t tank //we just declare the interface type
	//we need to assign it proper value
	fmt.Printf("%T",t)
	t = myValue{10,14}
	fmt.Printf("%T",t) //This gote the type now!!
	fmt.Println(t.area())
	fmt.Println(t.volume())
}

