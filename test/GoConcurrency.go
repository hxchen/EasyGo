package main

import (
	"fmt"
	"sync"
)

// MutexSync 互斥锁同步
func MutexSync() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("Hello, World!")
		mu.Unlock()
	}()

	mu.Lock()
}

// ChannelSync 通道同步
// 根据 Go 语言内存模型规范，对于从无缓冲 Channel 进行的接收，发生在对该 Channel 进行的发送完成之前。
// 因此，后台线程 <-done 接收操作完成之后，main 线程的 done <- 1 发送操作才可能完成（从而退出 main、退出程序），而此时打印工作已经完成了。
func ChannelSync() {
	done := make(chan bool)

	go func() {
		fmt.Println("Hello, World!")
		<-done
	}()
	// 无缓冲的通道只有在有人接收值的时候才能发送值。
	done <- true
}

func Recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func Send() {
	ch := make(chan int)
	go Recv(ch) // // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// ChannelCache 通道缓存 等待N个线程完成后的操作
func ChannelCache() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开 N 个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// 等待 N 个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

// WaitGroup 使用WaitGroup等待一组线程完成
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	wg.Wait()
}

func main() {
	//MutexSync()
	//ChannelSync()
	// Send()
	//ChannelCache()
	WaitGroup()
}
