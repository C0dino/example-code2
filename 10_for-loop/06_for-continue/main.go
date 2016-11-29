package main

import (
	"fmt"
)

func main() {
	i := 0 //init
	for {
		i++           //incrementer
		if i%2 == 0 { //condition
			continue
		}
		fmt.Println(i)
		if i >= 50 { //condition
			break
		}
	}
}

// This loop prints only odd numbers up to 51
