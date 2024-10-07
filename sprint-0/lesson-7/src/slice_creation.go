package main

import "fmt"

func main() {
	a := make([]int, 2)

	a[0] = 0
	a[1] = 1

	fmt.Println(a)

	var b []int
	fmt.Println(b == nil) // true

	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c) // [1 2 3 4 5]

	a1 := []int{1, 2, 3, 4, 5}
	b1 := a1[:4]
	c1 := b1[:3]

	fmt.Println(a1) // [1 2 3 4 5]
	fmt.Println(b1) // [1 2 3 4]
	fmt.Println(c1) // [1 2 3]
}
