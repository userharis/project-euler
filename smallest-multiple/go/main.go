package main

import "fmt"

func main() {
	fmt.Println(smallestMultiple(1, 20))
}

func smallestMultiple(min, max int) int {
	multiple := 1

	for ; min <= max; min++ {
		multiple = lcm(multiple, min)
	}

	return multiple
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	r := a % b

	if r == 0 {
		return b
	}

	return gcd(b, r)
}
