package main

import (
	"fmt"
	"math"
)

func discriminant(a, b, c float64) float64 {
	discriminant := math.Pow(b, 2) - 4*a*c

	return discriminant
}

func main() {
	fmt.Println(discriminant(1, -3.0, 2))
	fmt.Println(discriminant(1, 4.0, 2))
}
