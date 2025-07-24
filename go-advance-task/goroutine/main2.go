package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	taskList := make(chan func(), 5)
	go CreateTask(taskList, &wg)
	go ProcessTask(taskList, &wg)
	wg.Wait()
	close(taskList)
}

func CreateTask(tasklist chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// fmt.Printf("CreateTask start listLen: %v\n", len(tasklist))
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		fmt.Printf("[CreateTask][start],len: %v\n", len(tasklist))
		tasklist <- DoSomeThing
		fmt.Printf("[CreateTask][finish],len: %v\n", len(tasklist))
	}
}

// 怎么传入一个函数指针
func ProcessTask(taskList chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		fmt.Printf("[ProcessTask][start],taskNum = %v \n", len(taskList))
		task, ok := <-taskList
		fmt.Printf("[ProcessTask][finish],taskNum = %v \n", len(taskList))
		if !ok {
			return
		}
		startTime := time.Now()
		task()
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Printf("[ProcessTask][finish],duration: %v\n", elapsedTime)
	}
}

func DoSomeThing() {
	// 模拟一个随机耗时的任务
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Printf("DoSomeThing\n")
}
