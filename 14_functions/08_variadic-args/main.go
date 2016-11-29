package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} //slice of float64  Human example, this is "Broccoli"
	n := average(data...)                     // the ... in this case takes the values from data and sends them back out one at a time rather than all at once -- Makes Cupcakes
	fmt.Println(n)
}

func average(sf ...float64) float64 { //This is looking for "Cupcakes" -- says gtfo if you try to serve it "Broccoli"
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
