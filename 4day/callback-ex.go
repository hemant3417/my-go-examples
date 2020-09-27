package main

import (
	"fmt"
)

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(ii...)
	fmt.Println(total)

	evenTotal := evensum(sum, ii...)
	fmt.Println(evenTotal)
}
func evensum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
