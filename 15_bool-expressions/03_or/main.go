package main

import "fmt"

func main() {

	if true || false { // expressions happen horizontally
		fmt.Println("this ran")
	}
}
