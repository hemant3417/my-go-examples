package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg1 sync.WaitGroup

	wg1.Add(2)
	go func() {
		fmt.Println("Method 1.....")
		wg1.Done()
	}()
	go func() {
		fmt.Println("Method 2.....")
		wg1.Done()
	}()
	wg1.Wait()

}
