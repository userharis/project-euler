package main

import "fmt"

func main() {
	fmt.Println(sumOfEvenFibs(4000000))
}

func sumOfEvenFibs(n int) int {
	sum := 2

	if n < 2 {
		return 0
	}

	for f1, f2, f3 := 0, 2, 8; f3 <= n; f3 = 4*f2 + f1 {
		f1, f2 = f2, f3
		sum += f3
	}

	return sum
}
