package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isOperator(s string) bool{
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func check(question string) bool {
	questionArray := strings.Split(question," ")
	//fmt.Println(questionArray)
	for i:=0;i<len(questionArray);i++ {
		//fmt.Println(questionArray[i])
		if i%2 == 0 && isOperator(questionArray[i]) {
			return false
		} else if i%2 == 1 && !isOperator(questionArray[i]) {
			return false
		}
	}
	return true
}

func evaluate(question string) int {
	questionArray := strings.Split(question," ")
	//lenA := len(questionArray)
	operandStart := 2
	operatorStart := 1
	for ;operatorStart < len(questionArray);{
		op1, _ := strconv.Atoi(questionArray[0])
		op2, _ := strconv.Atoi(questionArray[operandStart])
		//fmt.Println(op1," ",op2)
		questionArray[operatorStart] = strings.TrimSpace(questionArray[operatorStart])
		//fmt.Println(questionArray[operatorStart])
		if questionArray[operatorStart] == "+" {
			//fmt.Println("jii")
			questionArray[0] = strconv.Itoa(op1 + op2)
		} else if questionArray[operatorStart] == "*" {
			questionArray[0] = strconv.Itoa(op1 * op2)
		} else if questionArray[operatorStart] == "/" {
			questionArray[0] = strconv.Itoa(op1 / op2)
		} else {
			questionArray[0] = strconv.Itoa(op1 - op2)
		}
		operandStart += 2
		operatorStart += 2
	}
	ans, _ := strconv.Atoi(questionArray[0])
	return ans
}

func Answer(question string) (int,bool){

	m := make(map[string]string)

	m["plus"] = "+"
	m["divided"] = "/"
	m["multiplied"] = "*"
	m["minus"] = "-"
	if question[0:8] != "What is " || question[len(question)-1] != '?' {
		return 0,false
	} else{
		//As all operators take just one word, we can replace multiplied by and divided by, we can just remove by
		question = strings.ReplaceAll(question,"divided by","divided")
		question = strings.ReplaceAll(question,"multiplied by","multiplied")

		questionArray := strings.Split(question," ")
		for i:=0;i<len(questionArray);i++ {
			//fmt.Println(questionArray[i])
			if val,ok := m[questionArray[i]]; ok {
				questionArray[i] = val
			}
		}
		index := len(questionArray)-1
		questionArray[len(questionArray)-1] = questionArray[index][0:len(questionArray[index])-1]
		//for i:=0;i<len(questionArray);i++ {
		//	fmt.Println(questionArray[i])
		//}
		majorQuestion := strings.Join(questionArray[2:]," ")
		if check(majorQuestion) {
			//Here, We have all the valid operations
			ans := evaluate(majorQuestion)
			fmt.Println(ans)
			return ans,true
		} else{
			return 0,false
		}
		//fmt.Println(majorQuestion)
	}
}

func main(){

	strings := []string{"What is 5?", "What is 1 plus 1?", "What is 53 plus 2?", "What is -1 plus -10?", "What is 123 plus 45678?","What is 4 minus -12?", "What is -3 multiplied by 25?", "What is 33 divided by -3?", "What is 1 plus 1 plus 1?","What is 1 plus 5 minus -2?","What is 20 minus 4 minus 13?", "What is 17 minus 6 plus 3?", "What is 2 multiplied by -2 multiplied by 3?", "What is -3 plus 7 multiplied by -2?", "What is -12 divided by 2 divided by -3?","What is 52 cubed?", "Who is the President of the United States?", "What is 1 plus?", "What is?", "What is 1 plus plus 2?","What is 1 plus 2 1?", "What is 1 2 plus?", "What is plus 1 2?",}

	for i:=0;i<len(strings);i++ {
		ans,_ := Answer(strings[i])
		fmt.Println(ans)
	}

}
