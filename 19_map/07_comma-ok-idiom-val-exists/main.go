package main

import (
	"fmt"
)

func main() {
	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[2]; exists {
		delete(myGreeting, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesnt exist.")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myGreeting)
}
