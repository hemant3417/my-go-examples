package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs : ", runtime.NumGoroutine())
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()

		}()
		fmt.Println("GoRoutines : ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())
	fmt.Println("counter : ", counter)

}
