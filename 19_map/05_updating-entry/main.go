package main

import (
	"fmt"
)

func main() {
	myGreeting := map[string]string{
		"Steve":   "Good Morning",
		"Jessica": "Bonjour",
	}
	myGreeting["Harley Quinn"] = "Howdy"
	fmt.Println(myGreeting)
	myGreeting["Harley Quinn"] = "Gidday"
	fmt.Println(myGreeting)
}
