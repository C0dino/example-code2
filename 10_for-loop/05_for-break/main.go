package main

import (
	"fmt"
)

func main() {
	i := 0 // init
	for {
		fmt.Println(i)
		if i >= 10 { // condition
			break
		}
		i++ //post
	}
}

// Set i := to 0 then run the loop until i is > or = to 10 then break
