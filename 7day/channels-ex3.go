package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	fmt.Println("-------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	fmt.Println("-------")
	fmt.Printf("\t%T\n", (<-chan int)(c))
	fmt.Printf("\t%T\n", (chan<- int)(c))

}
