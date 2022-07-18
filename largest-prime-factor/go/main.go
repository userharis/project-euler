package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(n int) int {
	sqrt := math.Sqrt(float64(n))
	var largestPrimeFactor int

	/*
	   since 2 is the only even prime number,
	   we run a separate loop for 2 and then
	   run another loop only for odd numbers
	*/

	if n%2 == 0 {
		largestPrimeFactor = 2
		// make n smaller, remove all factors that are multiples of 2
		for n%2 == 0 {
			n /= 2
		}
	}

	for d, s := 3, int(sqrt); d <= s; d += 2 {
		for n%d == 0 {
			largestPrimeFactor = d
			n /= d
		}
	}

	if largestPrimeFactor == 0 {
		// return n if its itself a prime number
		return n
	}
	return largestPrimeFactor
}
