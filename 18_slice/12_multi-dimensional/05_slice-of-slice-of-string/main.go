package main

import "fmt"

func main() {
	records := make([][]string, 0)
	// Student 1
	student1 := make([]string, 4)
	student1[0] = "Udemy"
	student1[1] = "Student"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	student2 := make([]string, 4)
	student2[0] = "Pendino"
	student2[1] = "Stephen"
	student2[2] = "50.00"
	student2[3] = "68.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)
}
