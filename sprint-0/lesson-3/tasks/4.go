package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Scan(&num1, &num2)

	if num1 > 0 && num2 > 0 {
		fmt.Println("Оба числа положительные")
	} else if num1 < 0 && num2 < 0 {
		fmt.Println("Оба числа отрицательные")
	} else if num1 == 0 || num2 == 0 {
		fmt.Println("Одно из чисел равно нулю")
	} else {
		fmt.Println("Одно число положительное, а другое отрицательное")
	}
}
