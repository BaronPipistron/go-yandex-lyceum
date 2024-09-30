package main

import "fmt"

func sumDigits(n int) int {
	var sum int = 0

	for n != 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func CalculateDigitalRoot(n int) int {
	for n > 10 {
		n = sumDigits(n)
	}

	return n
}

func main() {
	fmt.Println(CalculateDigitalRoot(12345))
}
