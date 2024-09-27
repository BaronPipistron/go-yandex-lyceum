package main

import "fmt"

func main() {
	var age int

	fmt.Println("Input your age: ")
	fmt.Scanln(&age)
	fmt.Println(age)
}
