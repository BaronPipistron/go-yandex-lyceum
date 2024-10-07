package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{4, 5, 2, 1, 3, 9, 7, 8, 6}
	fmt.Println(sort.IntsAreSorted(intSlice))
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	strSlice := []string{"Go", "Bravo", "Gopher", "YaLyceum", "Alpha", "Grin", "Delta"}
	sort.Strings(strSlice)
	fmt.Println(strSlice)
}
