package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var incrementor int64
	var wg sync.WaitGroup
	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			runtime.Gosched()
			fmt.Println("incrementor : ", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("end Value : ", incrementor)
}
