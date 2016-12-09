package main

import (
	"fmt"
	"runtime/internal/atomic"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intrn(3)) * time.Millisecond)
		//race:
		//counter++
		//no race:
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter: ", counter)
	}
	wg.Done()
}
