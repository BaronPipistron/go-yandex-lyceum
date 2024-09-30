package main

import "fmt"

func IsPowerOfTwoRecursive(n int) {
	if n == 2 || n == 1 {
		fmt.Println("YES")

		return
	} else if n%2 != 0 {
		fmt.Println("NO")

		return
	}

	IsPowerOfTwoRecursive(n / 2)
}

func main() {
	IsPowerOfTwoRecursive(5)
	IsPowerOfTwoRecursive(64)
	IsPowerOfTwoRecursive(128)
	IsPowerOfTwoRecursive(458)
}
