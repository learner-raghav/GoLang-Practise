package wordy

import (
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
	for i:=0;i<len(questionArray);i++ {
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

	operandStart := 2
	operatorStart := 1
	for ;operatorStart < len(questionArray);{
		op1, _ := strconv.Atoi(questionArray[0])
		op2, _ := strconv.Atoi(questionArray[operandStart])
		questionArray[operatorStart] = strings.TrimSpace(questionArray[operatorStart])
		if questionArray[operatorStart] == "+" {
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

		majorQuestion := strings.Join(questionArray[2:]," ")
		if check(majorQuestion) {
			//Here, We have all the valid operations
			ans := evaluate(majorQuestion)
			return ans,true
		} else{
			return 0,false
		}
	}
}