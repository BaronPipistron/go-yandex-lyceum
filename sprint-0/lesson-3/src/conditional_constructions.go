package main

import "fmt"

func main() {
	var score int

	fmt.Scanln(&score)

	if score >= 60 {
		fmt.Println("Nice work!")
	} else {
		fmt.Println("You need to study harder!")
	}
}
