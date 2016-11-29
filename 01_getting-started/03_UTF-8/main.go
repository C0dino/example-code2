package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i) // the \t stands for tabs | %d = Decimals | %b = Binary | %x = Hexadecimal | %q = UTF-8
	}
}
