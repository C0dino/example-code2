package main

import (
	"fmt"
)

func main() {

	b := true

	if food := "Chocolate"; b { //food := "Chocolate" is the init statement | b is the expression being checked by the if statement | This helps keep the scope tight
		fmt.Println(food)
	}
}
