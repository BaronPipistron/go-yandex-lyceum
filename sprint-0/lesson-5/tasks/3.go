package main

import "fmt"

func SumDigitsRecursive(n int) int {
	if n/10 == 0 {
		return n
	}

	return SumDigitsRecursive(n/10) + n%10
}

func main() {
	fmt.Println(SumDigitsRecursive(153))
}
