package main

import "fmt"

func main(){

	s:= []int{100,200,300,400}
	m := map[int]int {1:100,2:200,3:300,4:400}

	for _,j := range s{
		fmt.Print("( ",j,") ")
	}
	fmt.Println()
	for _,j := range m {
		fmt.Print("( ",j,") ")
	}
}