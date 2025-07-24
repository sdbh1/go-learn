// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。

package main

import (
	"fmt"
	"sync"
)

func main() {
	// 初始化计数器、互斥锁和等待组
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	// 启动 10 个协程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				// 加锁，保证同一时间只有一个协程能修改计数器
				mu.Lock()
				counter++
				// 解锁，允许其他协程修改计数器
				mu.Unlock()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出计数器最终值
	fmt.Printf("使用 sync.Mutex 实现的计数器最终值: %d\n", counter)
}
