package main

import (
	"fmt"
)

func max(x int) int {
    return 42 + x
    
}

func main()  {
    max := max (7)
    fmt.Println(max) //max is now the result, not the function (Lower case, only visible inside the package not outside)
}

// DONT DO THIS; bad coding practice!