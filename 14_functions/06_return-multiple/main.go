package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane ", "Doe ")) //spaces added after Jane & Doe to Separate them when they get returned below
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
