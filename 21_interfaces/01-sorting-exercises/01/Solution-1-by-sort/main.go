package main

import "sort"
import "fmt"

type people []string

func main() {
	studygroup := people{"zero", "john", "Al", "Jenny"}
	sort.Strings(studygroup)
	fmt.Println("Strings ", studygroup)
}
