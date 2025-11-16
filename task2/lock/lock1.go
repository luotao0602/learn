package main

import (
	"fmt"
	"sync"
)

func main() {
	//锁
	lock := sync.Mute{}
	var wg sync.WaitGroup
	count := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// 协程 -1
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// 加锁
				lock.Lock()
				count++
				// 释放锁
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count is: ", count)
}
