package main

import "fmt"

var fname string
var lname string

func main() {
	greet(&fname, &lname)
	fmt.Println("Hi!", fname+" "+lname)
}

func greet(fname, lname *string) {
	fmt.Print("Enter your first name: ") //using .Print because it keeps the text on the same line as the terminal cursor
	fmt.Scan(fname)                      //.Scan allows user via the terminal window
	fmt.Print("Enter your last name: ")

}

// This code is an exercise to show another way to use pointers in
// conjunction with multiple functions
