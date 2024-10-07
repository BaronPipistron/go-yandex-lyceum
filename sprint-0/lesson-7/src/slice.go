package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:]
	c := b[:]

	c[0] = 52

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
