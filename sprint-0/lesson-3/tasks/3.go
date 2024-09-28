package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	if number > 0 && number%2 == 0 {
		fmt.Println("Число положительное и четное")
	} else if number > 0 && number%2 != 0 {
		fmt.Println("Число положительное и нечетное")
	} else if number < 0 && number%2 == 0 {
		fmt.Println("Число отрицательное и четное")
	} else if number < 0 && number%2 != 0 {
		fmt.Println("Число отрицательное и нечетное")
	} else {
		fmt.Println("Число равно нулю")
	}
}
