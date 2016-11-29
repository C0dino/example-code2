package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i) // the \t stands for tabs | %d = Decimals | %b = Binary | %x = Hexadecimal
	}
}
