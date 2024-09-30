package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	var res int = 1
	for i := 2; i <= input; i++ {
		res *= i
	}

	fmt.Println(res)
}
