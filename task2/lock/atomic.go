package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var atomicNum int32
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// 协程 -1
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&atomicNum, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count is: ", atomicNum)
}
