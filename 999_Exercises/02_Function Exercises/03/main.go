package main

import "fmt"

func yeti(n ...int) int {
	var max int
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {

	greatest := yeti(25, 2, 3, 4, -3)
	fmt.Println(greatest)
}

/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers
*/
