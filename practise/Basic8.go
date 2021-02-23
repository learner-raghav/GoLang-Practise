package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X int
	Y int
}

func calculateDistance(v1 Vertex1, v2 Vertex1) float64 {
	ans := math.Sqrt(math.Pow(float64((v1.X-v2.X)),2) + math.Pow(float64((v1.Y-v2.Y)),2))
	return ans
}

func increaseByOne(s []int){
	fmt.Println(s)
}
// What is a Struct - A struct is a collection of fields.
func main() {

	//We might have points between which we want to calculate distance, in that case we can use struct

	v1 := Vertex1{2,3}
	v2 := Vertex1{6,6}


	ans := calculateDistance(v1,v2)
	fmt.Printf("The distance between the 2 points is %f \n",ans)


	pointerToV1 := &v1

	fmt.Println(pointerToV1)
	pointerToV1.Y = 23
	fmt.Println(*pointerToV1)
	fmt.Println(pointerToV1)

	array := [6]int{1,2,3,4,5,6}
	//for i:=0;i<6;i++ {
	//	fmt.Println(array[i])
	//}
	fmt.Println(array)

	s := array[4:5]
	//for i:=0;i<len(s);i++ {
	//	fmt.Println(s[i])
	//}
	fmt.Println(s)
	fmt.Println(cap(s)," ",len(s))

	fmt.Println(&s[0])
	fmt.Println(&array[1])

	var s1 []int
	fmt.Println(cap(s1)," ",len(s1))

	a := make([]int,5,10)
	fmt.Println(cap(a),len(a))
	b := a[3:5]
	fmt.Println(cap(b),len(b))


	// Understanding slices of slices
		board := [][]string {
			[]string{"_","_","_"},
			[]string{"_","_","_"},
			[]string{"_","_","_"},
		}
		for i:= range board{
			for j:= range board[0] {
				fmt.Print(board[i][j]+" ")
			}
			fmt.Println()
		}

		arr := []int {1,2,3,4,5}
		for i,j := range arr {
			fmt.Println(i,j)
		}

	var s3 []int
	fmt.Println(s3)
	s3 = append(s3,0,2,2)
	fmt.Println(s3)

	var arr3 = [8]int {6,5,4,3,2,1,0,-1}
	s5 := arr3[1:6]
	increaseByOne(s5)
	fmt.Println(s5)
	fmt.Println(arr3)
}
