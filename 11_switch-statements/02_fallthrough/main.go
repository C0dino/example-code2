package main

import (
	"fmt"
)

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Shushant")
	default:
		fmt.Println("You have no friends?")
	}
}

/*
no default fallthrough
fallthrough is optional
-- you can specify fallthrough by explicitly stating it
-- break isnt needed like in other languages
*/
