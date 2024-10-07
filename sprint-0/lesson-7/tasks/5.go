package main

import "fmt"

func Mix(nums []int) []int {
	newSlice := make([]int, len(nums))

	var xIndex int = 0
	var yIndex int = len(nums) / 2

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			newSlice[i] = nums[xIndex]
			xIndex++
		} else {
			newSlice[i] = nums[yIndex]
			yIndex++
		}
	}

	return newSlice
}

func main() {
	fmt.Println(Mix([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
