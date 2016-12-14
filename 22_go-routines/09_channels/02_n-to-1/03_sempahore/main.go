package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {

		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
