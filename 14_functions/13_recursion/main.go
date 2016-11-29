package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}

/* factorial is a sequence of numbers multiplied by the numbers beneath it.  In this example,
the factorial(4) is 4 * 3 * 2 * 1 * 1  which equals 24
Recursion is not typically the best solution.  Basically any solution that can be solved with recursion,
can be solved without recursion.  Recursion comes with a high cost of processing.  Most of the time it is
better to simply use loops.
*/
