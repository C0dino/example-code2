package main

import (
	"fmt"
)

func main() {
	fmt.Println(x)  //this does not work because x is declared after the println is called
	fmt.Println(y)  //this works because y is declared outside this code block at the package level
	x := 42
}

var y = 42

//this example again illustrates why order matters.
