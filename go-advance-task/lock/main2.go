// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	// 等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(10)
	fmt.Println("Before:", atomic.LoadInt64(&counter))

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("After:", atomic.LoadInt64(&counter))
}
