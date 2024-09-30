package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	var prev int = 0
	var curr int = 1

	for curr < input {
		var tmp int = curr
		curr += prev
		prev = tmp
	}

	for _ = range 10 {
		fmt.Println(curr)

		var tmp int = curr
		curr += prev
		prev = tmp
	}
}
