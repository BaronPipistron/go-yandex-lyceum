package main

import "fmt"

func ThirdElementInArray(array [6]int) int {
	return array[2]
}

func main() {
	var array [6]int = [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(ThirdElementInArray(array))
}
