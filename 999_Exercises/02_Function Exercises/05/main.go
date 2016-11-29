package main

import (
	"fmt"
)

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(x ...int) {
	fmt.Println(x)
}

/*
Write a function, foo, which can be called in all of these ways:
Func main() {
foo(1, 2)
foo(1, 2, 3)
aSlice := []int{1, 2, 3, 4}
foo(aSlice…)
foo()
}
*/
