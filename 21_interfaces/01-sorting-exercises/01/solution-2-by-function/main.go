package main

import (
	"fmt"
	"sort"
)

type people []string

func (s people) Len() int {
	return len(s)
}
func (s people) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s people) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func main() {

	studygroup := people{"zero", "john", "Al", "Jenny"}
	sort.Sort(people(studygroup))
	fmt.Println(studygroup)
}
