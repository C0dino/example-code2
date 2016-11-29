package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ { // init - i := 0 | condition i <= 5 | post - i++
		for j := 0; j <= 5; j++ { // init - i := 0 | condition i <= 5 | post - i++
			fmt.Println(i, " - ", j)
		}
	}
}

// Outer loop Runs, Gets to inner Loop, Inner Loop runs Until it reaches 5
// Then outer loop runs, gets to inner loop, inner loop runs until it reaches 5
// this repeats until the condition in the outer loop i <= 5
