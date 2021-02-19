package raindrops

import "strconv"

func Convert(input int) string{
	 ans := ""
	 val := false
	 if input%3 == 0 {
		 ans += "Pling"
		 val = true
	 }
	 if input%5 == 0 {
	 	ans += "Plang"
	 	val = true
	 }
	 if input%7 == 0 {
	 	ans += "Plong"
	 	val = true
	 }

	 if !val {
	 	ans += strconv.Itoa(input)
	 }
	 return ans
}