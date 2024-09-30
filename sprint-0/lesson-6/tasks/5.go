package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 7; i++ {
		fmt.Printf("%d я уже сделал: %s\n", i+1, array[i])
	}

	for i := 7; i < len(array); i++ {
		fmt.Printf("%d не успел сделать: %s\n", i+1, array[i])
	}
}

func main() {
	var array [9]string = [9]string{"a", "b", "c", "d", "e", "f", "g", "h", "k"}

	PrettyArrayOutput(array)
}
