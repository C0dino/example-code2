package main

import (
	"fmt"
)

func main() {
	i := 0        // init
	for i <= 10 { // condition
		fmt.Println(i)
		i++ // post
	}
}

// Different way of doing a condition
// loop will continue to iterate (run) until i <= 10
