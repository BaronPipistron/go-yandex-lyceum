package main

import "fmt"

func FiveSteps(array [5]int) [5]int {
	var ans [5]int
	var index int = 0

	for i := len(array) - 1; i >= 0; i-- {
		ans[index] = array[i]
		index++
	}

	return ans
}

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(FiveSteps(array))
}
