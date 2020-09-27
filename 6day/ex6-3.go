package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var incrementor int
	var wg sync.WaitGroup
	const gs = 100

	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incrementor
			runtime.Gosched()
			v++
			incrementor = v
			fmt.Println("incrementor : ", incrementor)
			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("end Value : ", incrementor)
}
