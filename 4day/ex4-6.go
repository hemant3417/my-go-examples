package main

import "fmt"

func main() {
	x := fool()
	y := fool()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(y())
	fmt.Println(y())
}

func fool() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
