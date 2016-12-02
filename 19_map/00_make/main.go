package main

import (
	"fmt"
)

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["John"] = "Good Morning"
	myGreeting["Jessica"] = "Bonjour."
	fmt.Println(myGreeting)
}
