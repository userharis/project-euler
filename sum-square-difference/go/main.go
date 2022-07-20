package main

import "fmt"

func main() {
	fmt.Println(sumSquareDifference(100))
}

func sumSquareDifference(n int) int {
	sumOfSquares := (n * (n + 1) * (2*n + 1)) / 6
	squareOfSum := (n * (n + 1) / 2) * (n * (n + 1) / 2)
	return squareOfSum - sumOfSquares
}
