package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi There")
	}()

	s := func(x int) {
		fmt.Println(x)
	}

	s(456)
}
