package main

import (
	"fmt"
)

func main() {
	var num1 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1) // Prompts user for input
	var num2 int
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2) // Prompts user for input
	answer := num1 / num2
	remainder := num1 % num2
	fmt.Println(num1, "divided by", num2, "equals", answer, "with a remainder of", remainder)
	// fmt.Println(num1, " divided by ", num2, "equals" answer, " remainder: ", remainder)
}
