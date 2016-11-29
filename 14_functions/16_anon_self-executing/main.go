package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I'm driving!")
	}()
}

/*
This is an anonymous self executing function.  It is NOT recommended to Code this way in go.
*/
