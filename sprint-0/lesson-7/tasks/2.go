package main

import (
	"fmt"
)

func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)

			i--
		}
	}

	return nums
}

func main() {
	fmt.Println(Clean([]int{6, 5, 4, 6, 2, 4, 6, 2, 1, 6}, 6))
}
