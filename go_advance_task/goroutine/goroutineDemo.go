package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func main() {
	// 记录 main 函数开始时间
	startTime := time.Now()

	ch := make(chan int)
	fmt.Printf("main print,hello world \n")
	var wg sync.WaitGroup
	wg.Add(2)

	go inputTask(ch, &wg)
	go printTask(ch, &wg)

	// 等待所有协程完成
	wg.Wait()
	// 关闭通道
	close(ch)
	fmt.Printf("main print,finish\n")

	// 记录 main 函数结束时间
	endTime := time.Now()
	// 计算运行时长
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("main 函数运行时长: %v\n", elapsedTime)
}

func inputTask(ch chan int, wg *sync.WaitGroup) {
	fmt.Printf("inputTask,start\n")
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	fmt.Printf("inputTask,finish\n")
}

func printTask(ch chan int, wg *sync.WaitGroup) {
	fmt.Printf("printTask,start\n")
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("printTask %v \n", <-ch)
	}
	fmt.Printf("printTask,finish\n")
}
