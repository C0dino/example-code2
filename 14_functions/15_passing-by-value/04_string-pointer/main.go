package main

import (
	"fmt"
)

func main() {
	name := "Steve"
	fmt.Println(&name) //Memory Address

	changeMe(&name)

	fmt.Println(&name) // Memory Address
	fmt.Println(name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
}
