package main

import (
	"fmt"
	"sync"
)

func main() {
	//定义通道
	count := 100
	ch := make(chan int, count)
	// 定义 WaitGroup，用于等待协程完成
	var wg sync.WaitGroup
	// 第一个协程往 channel 加数据
	wg.Add(1)
	go func() {
		//计数器 -1，协程已完成
		defer wg.Done()
		for i := 0; i <= count; i++ {
			ch <- i
		}
		// 必须关闭通道 否则会产生死锁
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println("接收到: ", num)
		}
	}()

	wg.Wait()
}
