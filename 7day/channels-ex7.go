package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("About to exit")

}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel ", v)

		case v := <-o:
			fmt.Println("From the odd channel ", v)

		case i, ok := <-q:
			if !ok {
				fmt.Println("From the quit channel ", i, ok)
				return
			} else {
				fmt.Println("From comma ok")
			}

		}
	}

}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
