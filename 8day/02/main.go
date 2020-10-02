package main

import "fmt"

func main() {

	is := []int{1, 5, 7, 9}
	fmt.Println(MySum(is...))

}

func MySum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
