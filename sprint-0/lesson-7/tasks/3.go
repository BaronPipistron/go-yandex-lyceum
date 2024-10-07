package main

import "fmt"

func SliceCopy(nums []int) []int {
	newSlice := make([]int, len(nums), len(nums))

	copy(newSlice, nums)

	return newSlice
}

func main() {
	fmt.Println(SliceCopy([]int{1, 2, 3, 4, 5}))
}
