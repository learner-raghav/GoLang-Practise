package main

import "fmt"

/**
	Multiple interfaces in Golang!!

	1. In Go language, the interface is a collection of method signatures and it is
		also a type means you can create a variable of an interface type.
 */
type AuthorDetails interface {
	details()
}

type AuthorArticles interface {
	articles()
}

type author struct {
	a_name, branch, college string
	year, salary, particles, tarticles int
}

func (a author) details(){
	fmt.Println("Name ",a.a_name)
	fmt.Println("Branch ",a.branch)
}

func (a author) articles(){
	fmt.Println("Pending articles: ",a.particles-a.tarticles)
}
func main(){
	values := author{
		"Mohit","ISE","MSRIT",2032,200000,100,60,
	}
	var i1 AuthorArticles
	i1 = values
	i1.articles()
	var i2 AuthorDetails
	i2 = values
	i2.details()
	fmt.Println()
}