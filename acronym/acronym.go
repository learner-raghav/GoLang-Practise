// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

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
	return strings.ToUpper(res)
}
