package main

import (
	"fmt"
)

func main() {
	x := 13 / 3
	fmt.Println("13 Divided by 3 is: ", x)
	z := 13 % 3
	fmt.Println("Remainder: ", z)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 70; i++ {
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("Even")
		}
	}
}
