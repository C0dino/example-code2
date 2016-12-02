package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Steve":   "Good Morning",
		"Jessica": "Bonjour",
	}
	fmt.Println(myGreeting["Jessica"])
}
