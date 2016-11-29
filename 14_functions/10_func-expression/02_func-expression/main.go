package main

import (
	"fmt"
)

func main() {
	greeting := func() { // considered an anonymous function because it has no identifier
		fmt.Println("Hello world!")
	}
	greeting()
}
