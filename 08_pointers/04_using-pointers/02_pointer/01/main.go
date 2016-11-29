package main

import (
	"fmt"
)

func zero(z *int) {
	*z = 0
}

func main() {
	x := 0
	zero(&x)
	fmt.Println(x) // x is 0
}
