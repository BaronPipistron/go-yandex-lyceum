package main

import "fmt"

func main() {
	var number int

	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Некорректный ввод")
	} else if number < 10 {
		fmt.Println("Число меньше 10")
	} else if number < 100 {
		fmt.Println("Число меньше 100")
	} else if number < 1000 {
		fmt.Println("Число меньше 1000")
	} else {
		fmt.Println("Число больше или равно 1000")
	}
}
