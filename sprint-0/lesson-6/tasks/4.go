package main

import "fmt"

func SumOfArray(array [6]int) int {
	var sum int = 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	return sum
}

func main() {
	var array [6]int = [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(SumOfArray(array))
}
