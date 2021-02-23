package main

import (
	"fmt"
	"strings"
)

func main(){
	s := "The Road _Not_ Taken"
	s = strings.ReplaceAll(s,"-"," ")
	s = strings.ReplaceAll(s,"_","")
	ans := strings.Fields(s)
	fmt.Println(ans)

	res := ""
	for _,j := range ans {
		//fmt.Println(j[0])
		res += string(j[0])
	}
	strings.ReplaceAll(res," ","")
	fmt.Println(strings.ToUpper(res))
}
