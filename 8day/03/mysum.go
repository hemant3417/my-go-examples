package main

import "fmt"

func main() {

	is := []int{1, 5, 7, 9}
	fmt.Println(mySum1(is...))

}

func mySum1(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
