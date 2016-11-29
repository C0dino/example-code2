package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("%d - %b - %#x \n", 42, 42, 42) // Adding the # front loads the zero | lower case x = All Lower Case
	// fmt.Printf("%d - %b - %#X \n", 42, 42, 42) // Upper Case X = All Capitals
	// fmt.Printf("%d - %b - %X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %X \n", 42, 42, 42) // the \t stands for tabs | %d = Decimals | %b = Binary | %x = Hexadecimal | %q = UTF-8
}
