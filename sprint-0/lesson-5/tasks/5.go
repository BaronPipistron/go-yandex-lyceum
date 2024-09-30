package main

import "fmt"

func CalculateSeriesSum(n int) float64 {
	var sum float64 = 1

	for i := 2; i <= n; i++ {
		sum += 1 / float64(i)
	}

	return sum
}

func main() {
	fmt.Println(CalculateSeriesSum(5))
}
