package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println(" About To Exit")

}

func foo(c chan<- int) {
	c <- 34
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
