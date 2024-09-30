package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
