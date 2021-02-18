// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func hasLetter(remark string) bool{

	for _,c := range remark {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
// Hey should have a comment documenting it.
func Hey(remark string) string {

	// removing the starting and trailing whitespaces
	remark  = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	//When string is in upper case and has a letter
	if strings.ToUpper(remark) == remark && hasLetter(remark) {
		if remark[len(remark)-1] == '?' {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else{
		if remark[len(remark)-1] == '?'{
			return "Sure."
		} else  {
			return "Whatever."
		}
	}
}
