package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {
	for i := range n {
		fmt.Printf("%d ", i+1)
	}
}

func main() {
	fmt.Println(Add(1, 2))
	fmt.Println(Multiply(2, 2))
	PrintNumbersAscending(10)
}
