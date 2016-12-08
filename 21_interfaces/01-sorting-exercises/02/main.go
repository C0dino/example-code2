package main

import (
	"fmt"
	"sort"
)

func main() {
	S := []string{"zero", "john", "Al", "Jenny"}
	sort.Strings(S)
	fmt.Println("Strings Sorted:", S)
}
