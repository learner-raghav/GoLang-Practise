package main

import "fmt"

//usual sum function!!
func sum(array []int) int {
	var sum1 int = 0
	for i:=0;i<len(array);i++ {
		sum1 += array[i]
	}
	return sum1
}

//This function is to calculate the factroial of a given number.
func factorial(n int) int {
	sum := 1
	for i:=1;i<=n;i++ {
		sum *= i
	}

	return sum
}

func main(){

	array := []int{1,2,3,4,5}
	for i:=0;i<len(array);i++ {
		fmt.Println(array[i])
	}

	//So, basically how do we declare an array in go is
	array1 := []string {"Raghav","is","a","boy"}

	//This is another way in which we work with loops.
	for index, element := range array1 {
		fmt.Println(index," ",element)
	}

	array3 := []int{1,2,3,4,5,6,7,8,9,10}
	ans := sum(array3)
	fmt.Println(ans)

	var n int

	fmt.Scanf("%d",&n)
	factorial := factorial(n)
	fmt.Printf("The factorial of %d is %d\n",n,factorial)
}