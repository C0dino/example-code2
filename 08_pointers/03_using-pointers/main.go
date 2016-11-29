package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // Memory Address (0x20818a220)

	var b *int = &a

	fmt.Println(b)  // Memory Address (0x20818a220)
	fmt.Println(*b) // 43

	*b = 42        // b says, "the value at this address, change it to 42""
	fmt.Println(a) //42

	//this is useful because we can pass a memory address instead of a bunch of values (we can pass a reference)
	//and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we dont have to pass around large amounts of data
	// we only have to pass around addresses

    // everything is PASS BY VALUE in go func
    // when we pass a memory address, we ARE passing a value    

}
