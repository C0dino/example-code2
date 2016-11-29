package main

import "fmt"

func main() {
	var y int
	fmt.Scan(y)
	fmt.Println(spression(y))
}

func spression(x int) (int, bool) {
	
	return x % 2, true
}

/*
write a function which takes an integer.  The function will have two returns.
The first return should be the argument divided by 2.  The second return should
be a bool that lets us know wether or not the argument was even. for example:
a. half(1) returns(0, false)
b. half(2) returns(1, true)
*/
