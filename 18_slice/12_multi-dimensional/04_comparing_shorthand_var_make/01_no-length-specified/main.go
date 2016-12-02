package main

import (
	"fmt"
)

func main() {
	student := []string{}    // length should be specified here
	students := [][]string{} // length should also be specified here
	student[0] = "Steve"
	// student = append(student, "Steve")
	fmt.Println(student)
	fmt.Println(students)
}

/*
Causes an error because no length of the array was specified
*/
