package main

import "fmt"

func main() {
	var num1, num2, num3 int

	n, err := fmt.Scan(&num1, &num2, &num3)

	if err != nil || n != 3 {
		fmt.Println("Некорректный ввод")

		return
	}

	if num1 == num2 && num1 == num3 {
		fmt.Println("Все числа равны")
	} else if num1 != num2 && num1 != num3 && num2 != num3 {
		fmt.Println("Все числа разные")
	} else {
		fmt.Println("Два числа равны")
	}
}
