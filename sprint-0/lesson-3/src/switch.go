package main

import "fmt"

func main() {
	var weather string

	fmt.Scanln(&weather)

	switch weather {
	case "rain":
		fmt.Println("Weather is rainy")
	case "sunny":
		fmt.Println("Weather is sunny")
	default:
		fmt.Println("Weather is not defined")
	}
}
