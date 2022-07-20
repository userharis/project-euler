package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(largestPalindromeProduct())
}

func largestPalindromeProduct() int {
	result := 0

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			product := i * j
			s := strconv.Itoa(product)
			r := reverse(s)

			if r == s && product > result {
				result = product
			}
		}
	}

	return result
}

func reverse(str string) string {
	reversed := ""
	for _, r := range str {
		reversed = string(r) + reversed
	}
	return reversed
}
