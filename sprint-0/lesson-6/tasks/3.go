package main

import "fmt"

func FindMinMaxInArray(array [10]int) (int, int) {
	var max int = array[0]
	var min int = array[0]

	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}

		if array[i] < min {
			min = array[i]
		}
	}

	return max, min
}

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var max, min int = FindMinMaxInArray(array)

	fmt.Println(max, min)
}
