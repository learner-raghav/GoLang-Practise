package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {

	if str == "" {
		return ""
	}

	var result string = ""
	fmt.Println("hello")
	for i:= len(str)-1;i>=0;i-- {
		result += string(str[i])
		fmt.Println(result+"hi")
	}

	return result
}

func main(){

	var name string = "Raghav"
	fmt.Println(name)
	for i:=0 ; i< len(name); i++ {
		fmt.Println(string(name[i]))
	}

	a := 2
	b := 3
	fmt.Println(a," ",b)

	a,b = b,a
	fmt.Println(a," ",b)

	var fName string = "hello world"
	var reverseString string = ""

	for i:=len(fName)-1;i>=0;i-- {
		reverseString += string(fName[i])
	}

	fmt.Println("Original String: ", fName)
	fmt.Println("Modified String: ", reverseString)
	string1 := "nitin"
	string2 := Reverse(string1)

	fmt.Println(string2)

	str := "Raghav"
	charArray := strings.Split(str,"")
	for i:=0;i<len(charArray);i++ {
		fmt.Println(charArray[i])
	}

	for i:=0;i<len(charArray)/2;i++ {
		charArray[i],charArray[len(charArray)-i-1] = charArray[len(charArray)-i-1],charArray[i]
 	}
 	ans := strings.Join(charArray,"")
 	fmt.Println(ans)


 	string3 := " Raghav "
 	fmt.Println(string3)

 	string3 = strings.Trim(string3," ")

 	fmt.Println("hi"+string3+"hi")

 	tempStr := "HELLO WORLD"
 	if strings.ToUpper(tempStr) == tempStr {
 		fmt.Println("All in upper case")
	}

	tempStr1 := "hello world"
	if strings.ToLower(tempStr1) == tempStr1 {
		fmt.Println("All in lower case")
	}
}
