package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPUs : ", runtime.NumGoroutine())
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
			mu.Unlock()
		}()
		fmt.Println("GoRoutines : ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)

}
