package main

import "fmt"

func main() {
	defer last()
	xiz := []int{1, 2, 4, 5}
	fmt.Println(foo1(xiz...))
	fmt.Println(bar1(xiz))

}

func foo1(x ...int) int {
	defer func() {
		fmt.Println("Defer in fooo")
	}()

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("Foo ran")
	return sum
}

func bar1(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func last() {
	fmt.Println("Last line of exec")
}
