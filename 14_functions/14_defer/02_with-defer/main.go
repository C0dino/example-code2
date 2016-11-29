package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Println("World")
}
func main() {
	defer world() // world() will be defered until the rest of the main func runs (Right before func main exits)
	hello()
}

/* defer is typically used to defer closing a filepath
 */
