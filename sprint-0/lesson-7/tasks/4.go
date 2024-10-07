package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	newSlice := make([]int, len(nums1)+len(nums2), len(nums1)+len(nums2))

	for i, num := range nums1 {
		newSlice[i] = num
	}

	for i, num := range nums2 {
		newSlice[i+len(nums1)] = num
	}

	return newSlice
}

func main() {
	fmt.Println(Join([]int{1, 2, 3}, []int{1, 2, 3}))
}
