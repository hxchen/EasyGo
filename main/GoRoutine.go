package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d, s == %s\n", i, s)
	}
}

// sync.WaitGroup 用于等待 say("world") 完成。
// wg.Add(1) 增加了一个等待计数，
// defer wg.Done() 确保在 say("world") 结束时减少计数，
// wg.Wait() 等待所有的 wg.Done() 调用完成，然后主 goroutine 才会结束。
// 这样可以确保 say("world") 在主 goroutine 退出之前完成其所有输出
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say("world", &wg)
	say("hello", &wg)
	wg.Wait()
}
