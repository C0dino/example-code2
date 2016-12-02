package main

import (
	"fmt"
)

func main() {
	myGreeting := make(map[string]string)
	myGreeting["John"] = "Good Morning"
	myGreeting["Jessica"] = "Bonjour."
	fmt.Println(myGreeting)
}
