package main

import (
	"fmt"
	"time"
)

// Go 并发编程之 channel
// 一个通道发送和接收数据，默认是阻塞的。
// 当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个 Goroutine 从该通道读取数据。
// 相对地，当从通道读取数据时，读取被阻塞，直到一个 Goroutine 将数据写入该通道。
// 这些通道的特性是帮助 Goroutines 有效地进行通信，而无需像使用其他编程语言中非常常见的显式锁或条件变量。

func channelTest() {

	// 创建一个 channel
	ch1 := make(chan bool)
	fmt.Println("ch1:", ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子 Goroutine 中，i:", i)
		}
		// 循环结束后，向 channel 中写数据，表示要结束了
		ch1 <- true

		fmt.Println("子 Goroutine 结束...")
	}()

	// 从 ch1 通道中读取数据
	// 再子 Goroutine 写数据之前，主 Goroutine 是会被阻塞的
	fmt.Println("主 Goroutine 被阻塞..")
	data := <-ch1
	fmt.Println("主 Goroutine 拿到数据了，data:", data)

}

// Go 并发编程之 channel 缓冲区
// 可以通过设置缓冲区来避免阻塞
func channelBuffSizeTest() {

	// 创建一个 channel，待缓冲区，写满 5 个之前不会阻塞
	ch1 := make(chan int, 5)
	fmt.Println("ch1:", ch1)

	go func() {
		time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			data := <-ch1
			fmt.Println("子 Goroutine 拿数据，i:", data)
		}
		fmt.Println("子 Goroutine 结束...")
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("主 Goroutine 写数据，i:", i)
		ch1 <- i
	}

	fmt.Println("主 Goroutine 结束...")

	time.Sleep(2 * time.Second)

}

func main() {
	//channelTest()
	channelBuffSizeTest()
}
