package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57) //type is not yet determined they are "of a kind" --- these are literal values, constants -- example of the parallel constant type system
	fmt.Println(n)
}

func average(sf ...float64) float64 { //prefix ... allows for passing in as many float64's (in this example) as you want
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64      //total is an accumulator - a place where you accumulate values
	for _, v := range sf { // range can be used to loop over something -- the numbers entered into average above in this case
		total += v
	}
	return total / float64(len(sf))
}
