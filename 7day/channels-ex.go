package main

import "fmt"

func main() {
	v := make(chan int)
	go func() {
		v <- 22
	}()
	//fmt.Println(<-v)

	val, ok := <-v
	fmt.Println(val, ok)

	val, ok := <-v
	fmt.Println(val, ok)

}
