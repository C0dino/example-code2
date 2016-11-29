package main

import (
	"fmt"
)

func main() { // greet being declared as a parameter
	greet("Jane")
	greet("John")
}

func greet(name string) { // parameter name of type string
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
