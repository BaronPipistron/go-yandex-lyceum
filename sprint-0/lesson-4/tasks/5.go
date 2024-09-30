package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	if input < 0 {
		fmt.Println("Некорректный ввод")

		return
	}

	var sum int = 0
	for i := 1; i <= input; i += 2 {
		sum += i
	}

	fmt.Println(sum)
}
