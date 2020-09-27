package main

import "fmt"

func main() {
	i := foo()
	fmt.Println(i)
	a, b := bar()
	fmt.Println(a, b)
}

func foo() int {
	return 453
}

func bar() (int, string) {
	return 35, "Hello there"
}
