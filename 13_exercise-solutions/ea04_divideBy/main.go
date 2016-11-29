package main

import "fmt"

func main() {
	fmt.Print("What is the first number?: ")
	var num1 int
	fmt.Scan(&num1)
	fmt.Print("What is the second number?: ")
	var num2 int
	fmt.Scan(&num2)
	println(num1, "divided by", num2, "equals:", num1/num2, "With a remainder of:", num1%num2)
}
