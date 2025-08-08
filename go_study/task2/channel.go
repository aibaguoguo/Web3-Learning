package main

import (
	"fmt"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
func testChanel1() {
	ch := make(chan int)
	defer close(ch)
	//发
	go func(ch chan<- int) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}(ch)

	//收
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)

	time.Sleep(time.Second)
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func testChanel2() {
	ch := make(chan int, 10)
	defer close(ch)
	//发
	go func(ch chan<- int) {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}(ch)

	//收
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
	}(ch)

	time.Sleep(time.Second)
}
