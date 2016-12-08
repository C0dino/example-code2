package main

import "sort"
import "fmt"

type people []string

func main() {
	studygroup := people{"zero", "john", "Al", "Jenny"}
	//sort.Strings(studygroup)
	// sort.Sort(sort.StringSlice(studygroup))
	sort.StringSlice(studygroup).Sort()
	fmt.Println("Strings ", studygroup)
}
