package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

}

// 题目1:编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printNums() {
	//打印奇数
	go func() {
		for i := 1; i < 10; i += 2 {
			fmt.Printf("打印奇数:%d\n", i)
		}
	}()

	//打印偶数
	go func() {
		for i := 0; i <= 10; i += 2 {
			fmt.Printf("打印偶数:%d\n", i)
		}
	}()

	time.Sleep(2 * time.Second)
}

//题目2:设计一个任务调度器，接收一组任务（可以用函数表示），
//并使用协程并发执行这些任务，同时统计每个任务的执行时间。

func doJob(name string, task func()) {
	now := time.Now().UnixMilli()
	task()
	fmt.Println("任务"+name+"当前任务耗时", time.Now().UnixMilli()-now)
}
