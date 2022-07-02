package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(1000))
}

func sumOfMultiples(n int) int {
	sum := 0

	for i := 3; i < n; i += 3 {
		sum += i
	}

	for i := 5; i < n; i += 5 {
		sum += i
	}

	return sum
}
