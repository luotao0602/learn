package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义任务结构体
type Task struct {
	name string       //任务名
	Func func() error // 任务函数：返回 error 表示执行失败
}

type TaskResult struct {
	taskName string        //任务名称
	costTime time.Duration //任务耗时
	Err      error         // 执行错误（nil 表示成功）
}

// 定义任务调度器结构体
type Scheduler struct {
	maxWorkers     int             // 最大并发数
	taskChan       chan Task       // 任务通道
	taskResultChan chan TaskResult // 结果通道
	wg             sync.WaitGroup  // 等待所有任务完成
}

func newScheduler(maxTasks int) *Scheduler {
	return &Scheduler{
		maxWorkers:     maxTasks,                        // 最大并发数赋值
		taskChan:       make(chan Task, maxTasks),       // 缓冲通道大小 = 最大并发数
		taskResultChan: make(chan TaskResult, maxTasks), // 结果通道（缓冲足够大，避免阻塞）
	}
}

// 提交任务
func (s *Scheduler) submit(task Task) {
	s.taskChan <- task // 将任务发送到任务通道（缓冲满时会阻塞，控制并发）
	s.wg.Add(1)        //任务计数+1
}

// Close 关闭任务通道（所有任务提交完成后调用）
func (s *Scheduler) closeChan() {
	close(s.taskChan)
}

// 返回任务执行结果
func (s *Scheduler) returnResult() <-chan TaskResult {
	return s.taskResultChan
}

// executeTask 执行单个任务，统计执行时间并收集结果
func (s *Scheduler) execute(task Task) {
	startTime := time.Now() // 记录任务开始时间
	Error := task.Func()
	costTime := time.Since(startTime) // 计算耗时

	s.taskResultChan <- TaskResult{
		taskName: task.name,
		costTime: costTime,
		Err:      Error,
	}
}

// executeTask 执行单个任务，统计执行时间并收集结果
func (s *Scheduler) start() {
	for i := 0; i < s.maxWorkers; i++ {
		go func(workerID int) {
			// 循环读取任务通道，直到通道关闭
			for task := range s.taskChan {
				// 修复循环变量捕获：将 task 作为参数传入匿名函数
				func(t Task) {
					defer s.wg.Done() // 确保任务执行完后，计数-1（即使 panic 也执行）
					s.execute(t)
				}(task) // 传入当前循环的 task 副本
			}
		}(i)
	}
}

// 等待任务执行完
func (s *Scheduler) wait() {
	go func() {
		s.wg.Wait()
		// 不关闭 会报死锁
		close(s.taskResultChan) // 所有任务处理完成，关闭任务结果通道,
	}()
}

func buildTasks() []Task {
	return []Task{
		{
			name: "任务1-快速任务",
			Func: func() error {
				time.Sleep(500 * time.Millisecond) // 模拟 500ms 耗时
				return nil                         // 执行成功
			},
		},
		{
			name: "任务2-中等耗时",
			Func: func() error {
				time.Sleep(1 * time.Second) // 模拟 1s 耗时
				return nil
			},
		},
		{
			name: "任务3-失败任务",
			Func: func() error {
				time.Sleep(300 * time.Millisecond)
				return fmt.Errorf("数据库连接超时") // 模拟执行失败
			},
		},
		{
			name: "任务4-长耗时任务",
			Func: func() error {
				time.Sleep(2 * time.Second) // 模拟 2s 耗时
				return nil
			},
		},
		{
			name: "任务5-快速任务",
			Func: func() error {
				time.Sleep(400 * time.Millisecond)
				return nil
			},
		},
	}
}

func main() {
	scheduler := newScheduler(6)

	tasks := buildTasks()

	// 4. 提交所有任务到调度器
	go func() {
		for _, task := range tasks {
			scheduler.submit(task)
			fmt.Printf("已提交任务：%s\n", task.name)
		}
		scheduler.closeChan() // 所有任务提交完成，关闭任务通道
	}()

	//执行任务
	scheduler.start()

	//等待任务执行完
	scheduler.wait()

	// 5. 遍历结果通道，输出每个任务的执行情况
	fmt.Println("\n===== 任务执行结果 =====")

	for result := range scheduler.returnResult() {
		status := "成功"
		if result.Err != nil {
			status = "失败"
		}
		fmt.Printf("任务：%s | 状态：%s | 耗时：%v | 错误：%v\n",
			result.taskName, status, result.costTime.Round(10*time.Millisecond), result.Err)
	}
}
