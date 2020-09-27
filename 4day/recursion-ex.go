package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println(f)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
