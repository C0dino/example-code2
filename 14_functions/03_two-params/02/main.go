package main

import (
	"fmt"
)

func main() {  // calling function
	greet("Jane", "Doe") // passing in two arguments
}

func greet(fname, lname string) { //function declared with two parameters -- First Name & Last Name
	fmt.Println(fname, lname)
}
