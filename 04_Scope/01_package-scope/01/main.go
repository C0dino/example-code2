package main

import (
	"fmt"
)

// var x declares the variable | int is the TYPE of variable | 42 is being assigned to the variable which initializes the variable
// the scope of x is the entire package because it is accesible everywhere within this package
var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
	// Declaring y := 17 after the print line will not work, order matters
}

func foo() {
	fmt.Println(x)
	//	fmt.Println(y)
	// y is not available in this code block because it was defined in the block above, thats considered block level
}
