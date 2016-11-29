package main

import (
	"fmt"
	"crypto/aes"
)

func main()  {
    
    a := 43

    fmt.Println(a)
    fmt.Println(&a)

    var b *int = &a
    
    fmt.Println(b)

    /* the above code makes b a pointer to the memory address where an int is stored
    b is of type "int pointer"
    *int -- the * is part of the type -- b is of type *int
    * means pointer, * is POINTING to the int stored at the memory location of &a
    */
}