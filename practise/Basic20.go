package main

import "fmt"

func main(){
	m := map[string]string{}

	m["name"] = "Raghav"
	m["education"] = "B.E.(Pursuing)"

	fmt.Println(m)
	//This implies maps are passed by reference by default.
	func1(m)
	fmt.Println(m)
}

func func1(m map[string]string){
	m["life"] = "chill"
	m["name"] = "hola"
}