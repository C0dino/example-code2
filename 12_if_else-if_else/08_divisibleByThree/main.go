package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Prints out any number divisible by 3 up to 100
