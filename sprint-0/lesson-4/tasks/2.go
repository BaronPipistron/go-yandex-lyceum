package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	for i := range 10 {
		fmt.Printf("%d * %d = %d\n", input, i+1, input*(i+1))
	}
}
