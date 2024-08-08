package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 全局变量
var ticket = 10 //10 张票

// 锁
var waitGroup sync.WaitGroup
var mutex sync.Mutex

func main() {
	waitGroup.Add(4)
	// 4 个 Goroutine，代表 4 个售票窗口，它们会操作同一个变量 ticket
	go saleTicket("售票口1")
	go saleTicket("售票口2")
	go saleTicket("售票口3")
	go saleTicket("售票口4")

	// 阻塞主 Goroutine，等待 4 个协程调用 waitGroup.Done()
	waitGroup.Wait()
	fmt.Println("售票结束...")
}

func saleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	// 调用 waitGroup.Done()，4 个协程都调用后主 Goroutine 就可以继续执行了
	defer waitGroup.Done()
	for {
		// 加锁
		mutex.Lock()
		if ticket > 0 {
			// 睡眠，增大错误发生的概率
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// 售票
			fmt.Println(name, "售出：", ticket)
			// 减票
			ticket--
		} else {
			// 解锁 不可漏掉
			mutex.Unlock()
			fmt.Println(name, "售罄，没票了...")
			break
		}
		// 解锁
		mutex.Unlock()
	}
}
