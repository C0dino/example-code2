package main

import "fmt"

func main() {
	var fac, cardCode int32 = 2564, 889
	fac = fac*padding(cardCode) + cardCode
	fmt.Println(fac)
}

func padding(n int32) int32 {
	var p int32 = 1
	for p < n {
		p *= 10
	}
	return p
}
