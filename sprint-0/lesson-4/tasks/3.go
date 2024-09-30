package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	// sum := (1 + input) * input / 2

	var sum int = 0
	for i := range input {
		sum += (i + 1)
	}

	fmt.Println(sum)
}
