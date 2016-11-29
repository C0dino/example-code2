package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Print("Enter your name: ") //using Print keeps the entry on the same line as the scan.  If you were to use PrintLn it would end up underneath the line stating "Enter Your Name: "
	fmt.Scan(&name)
	switch name {
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
