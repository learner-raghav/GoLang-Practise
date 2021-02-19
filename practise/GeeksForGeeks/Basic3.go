package main

import "fmt"

/**

	A student program in golang

	1. We make a student structure where each student will have the resume.
 */

func main(){

	type Resume struct {
		college,city,department string
		cgpa float32
		skills []string
	}

	type Student struct {
		name, fatherName string
		resume Resume
	}

	student1 := Student{name: "raghav",fatherName: "keshav",resume: Resume{
		college: "ms ramaiah institute of technology",
		cgpa : 9.0,
		department: "ISE",
		skills: []string{"java","python","c"},
	}}

	fmt.Println(student1)

	var students [10]Student

	for i := range students {
		students[i] = student1
	}
	fmt.Println(students)
}