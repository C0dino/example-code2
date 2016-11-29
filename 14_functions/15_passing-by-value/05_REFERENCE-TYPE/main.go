package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Steve"]) // 32
}

func changeMe(z map[string]int) {
	z["Steve"] = 32
}

/*
Allocation with makeback to allocation. the built-in function make (T, args)
serves a purpose different from new(T). It creates slices, maps, and channels only,
and it returns an initialized (not zeroed) value of type T (not *T). the reason for
the distinction is that these three types represent, under the covers, references to data structures
that must be initialized before use.  A slice, for example, is a three-item descriptor containing
a pointer to the data (inside an array), the length, and the capacity, and until those items
the slice is nil.  For slices, maps, and channels, make initializes the interneal data structure
*/
