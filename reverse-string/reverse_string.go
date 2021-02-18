package reverse

import (
	"strings"
)

/**
	Good Algorithm:
		str = "raghav"
		i = 0, interchange with len(str) - i - 1
		Repeat the above steps till half the length of the string.

 */
func Reverse(str string) string {

	if str == "" {
		return ""
	}

	charArray := strings.Split(str,"") //splitting at each character

	for i:=0;i<len(charArray)/2;i++ {
		//swapping the characters
		charArray[i],charArray[len(charArray)-i-1] = charArray[len(charArray)-i-1],charArray[i]
	}

	//combining the individual strings
	result := strings.Join(charArray,"")
	return result
}