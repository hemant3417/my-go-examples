package main

import "fmt"

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println(evensum1(sum1, arr1...))

}

func evensum1(f func(xi ...int) int, yi ...int) int {
	var ei []int
	for _, v := range yi {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}
	return f(ei...)
}

func sum1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
