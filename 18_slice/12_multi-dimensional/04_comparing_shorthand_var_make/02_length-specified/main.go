package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)    // length should be specified here
	students := make([][]string, 35) // length should also be specified here
	student[0] = "Steve"
	// student = append(student, "Steve Squared")
	fmt.Println(student)
	fmt.Println(students)
}

/*
Causes an error because no length of the array was specified
*/
