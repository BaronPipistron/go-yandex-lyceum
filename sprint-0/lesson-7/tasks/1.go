package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil || n < 0 {
		return nil, errors.New("incorrect input")
	}

	var ans []int

	for _, num := range nums {
		if len(ans) == n {
			break
		}

		if num < limit {
			ans = append(ans, num)
		}
	}

	return ans, nil
}

func main() {
	test := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}

	fmt.Println(UnderLimit(test, 3, 5))
}
