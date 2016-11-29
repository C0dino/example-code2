package main

import "fmt"

func main() {

	x := []int{25, 2, 3, 4, -3}
	yeti(x...)
	fmt.Println(x)
}

func yeti(n ...int) {
	var max, min int
	for _, v := range n {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
		fmt.Println(max, min)
	}
}

/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers
*/
