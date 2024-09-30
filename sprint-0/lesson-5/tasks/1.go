package main

import (
	"fmt"
	"math"
)

func calculateDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func SqRoots() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	discriminant := calculateDiscriminant(a, b, c)

	if discriminant == 0 {
		fmt.Println(-b / (2 * a))

		return
	} else if discriminant < 0 {
		fmt.Println(0, 0)

		return
	}

	x_1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	x_2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	fmt.Println(x_1, x_2)
}

func main() {
	SqRoots()
}
