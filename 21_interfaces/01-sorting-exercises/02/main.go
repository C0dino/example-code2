package main

import (
	"fmt"
	"sort"
)

func main() {
	S := []string{"zero", "john", "Al", "Jenny"}
	sort.Strings(S)
	fmt.Println("Strings Sorted:", S)
	sort.Sort(sort.Reverse(sort.StringSlice(S)))
	fmt.Println("Strings Reversed: ", S)
}
