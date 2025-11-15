package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义 WaitGroup，用于等待两个协程完成
	var wg sync.WaitGroup

	// 计数器加 1，表示要等待一个协程
	wg.Add(1)
	go func() {
		defer wg.Done() // 协程结束时，计数器减 1
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	// 主协程等待所有工作协程完成（计数器变为 0）
	wg.Wait()
	fmt.Println("所有协程执行完毕")
}
