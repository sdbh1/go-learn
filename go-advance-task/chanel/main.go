package main

import (
	"fmt"
	"os"
	"strconv"
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

	var cachaNum int = 0

	if len(os.Args) > 1 {
		var err error
		cachaNum, err = strconv.Atoi(os.Args[1])
		if err != nil {
			cachaNum = 0
		}
	}

	StartChanelCacha(cachaNum)

	// 记录 main 函数结束时间
	endTime := time.Now()
	// 计算运行时长
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("main 函数运行时长: %v\n", elapsedTime)
}

func StartChanelCacha(cachaNum int) {
	fmt.Printf("StartChanelCacha %v \n", cachaNum)
	var ch chan int
	if cachaNum == 0 {
		ch = make(chan int)

	} else {
		ch = make(chan int, cachaNum)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// 启动生产和消费协程
	go produce(ch, &wg)
	go consume(ch, &wg)

	// 等待所有协程完成
	wg.Wait()
}

// produce 函数用于生产数据并发送到通道中
func produce(ch chan int, wg *sync.WaitGroup) {
	fmt.Printf("produceTask,start\n")
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		ch <- i
	}
	// 关闭通道
	close(ch)
	fmt.Printf("produceTask,finish\n")
}

// consume 函数用于从通道中接收数据并打印
func consume(ch chan int, wg *sync.WaitGroup) {
	fmt.Printf("consumeTask,start\n")
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		<-ch
	}
	fmt.Printf("consumeTask,finish\n")
}
