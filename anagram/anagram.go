package anagram

import (
	"strings"
)

//Not optimized!!

func sortString(array []byte,n int) string{

	for i:=0;i<=n-1;i++ {
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

func Detect(str string, candidates []string) []string{


	sortedWord := sortString([]byte(strings.ToLower(str)),len(str))
	var res []string

	//Iterating over the candidates and checking which are anagrams
	for i:=0;i<len(candidates);i++ {

		candidate := strings.ToLower(candidates[i])
		if candidate == strings.ToLower(str) {
			continue
		}
		sortedStr := sortString([]byte(candidate),len(candidate))
		if sortedStr == sortedWord {
			res = append(res,candidates[i])
		}
	}
	return res
}