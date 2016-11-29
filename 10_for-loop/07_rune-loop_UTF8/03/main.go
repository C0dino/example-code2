package main

import (
	"fmt"
)

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Println('i', "-", string(i), "-", []byte(string(i)))  // using the single 'i' uses the rune of i and references 105 the entire time rather than counting
	}
	foo := 'a'  // storing foo in a 'rune' as 'a' then printing out foo (on the next line) shows the value for a (97)
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T\n", rune(foo))
}
