package main

import "fmt"

//check if a given string is an anagram or not

func checkAnagram(str1 string, str2 string) bool{
	return true
}


func sort(array [6]int) [6]int{

	//Using normal bubble sort
	n := len(array)
	for i:=0;i<=n-1;i++ {
		fmt.Println(array)
		for j := 0; j <n-1-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}

func sortString(array []byte,n int) string{

	for i:=0;i<=n-1;i++ {
		fmt.Println(array)
		for j := 0; j <n-1-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j+1]
				array[j+1] = array[j]
				array[j] = temp
			}
		}
	}
	sortedString := string(array)
	return sortedString
}

func main(){

	var arr = [6]int {6,5,4,3,2,1}

	arr = sort(arr)

	for i:=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}

	str := "Raghav"
	var charArray []byte
	charArray = []byte(str)

	sortedStr1 := sortString(charArray,len(str))
	fmt.Println(sortedStr1)



}