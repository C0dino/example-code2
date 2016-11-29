package main

import (
	"fmt"
)

func main() {
	i := 0 // init
	for {
		fmt.Println(i)
		i++ //post
	}
}

//No condition, will run forever or until Stopped (Ctrl + C Twice)
