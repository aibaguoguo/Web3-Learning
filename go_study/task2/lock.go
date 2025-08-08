package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
type SafeCounter struct {
	count int
	lock  sync.Mutex
}

// 增加计数器
func (counter *SafeCounter) incr() int {
	counter.lock.Lock()
	defer counter.lock.Unlock()
	counter.count++
	return counter.count
}

func testLock1() {
	c := SafeCounter{}
	for i := 0; i < 10; i++ {
		go func(c *SafeCounter, gorId int) {
			for j := 0; j < 1000; j++ {
				fmt.Println("协程ID", i, "计数器的值:", c.incr())
			}
		}(&c, i)
	}
	time.Sleep(time.Second)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
type AtomicCounter struct {
	count int32
}

func (counter *AtomicCounter) incr() int32 {
	return atomic.AddInt32(&counter.count, 1)
}

func main() {
	c := AtomicCounter{}
	for i := 0; i < 10; i++ {
		go func(c *AtomicCounter, gorId int) {
			for j := 0; j < 1000; j++ {
				fmt.Println("协程ID", i, "计数器的值:", c.incr())
			}
		}(&c, i)
	}
	time.Sleep(time.Second)
}
