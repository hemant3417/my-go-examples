package main

import "fmt"

func main() {
	re := factorial1(3)

	fmt.Println(re)

}

func factorial1(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial1(x-1)
}
