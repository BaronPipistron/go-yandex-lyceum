package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	for i := 3; i <= input; i += 3 {
		fmt.Printf("%d\n", i)
	}
}
