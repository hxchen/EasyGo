package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main Go 语言中，要传递某个数据给另一个 Goroutine，可以把这个数据封装成一个对象，
// 然后把这个对象的指针传入某个 channel 中，另外一个 Goroutine 从这个 channel 中读出这个指针，并处理其指向的内存对象。
// Go 从语言层面保证同一个时间只有一个 Goroutine 能够访问 channel 里面的数据，为开发者提供了一种优雅简单的工具，
// 所以 Go 的做法就是使用 channel 来通信，通过通信来传递内存数据，使得内存数据在不同的 Goroutine 中传递，而不是使用共享内存来通信。
func main() {
	var ticket = 10
	var ticketChan = make(chan int)

	go sellTicket("售票口1", ticketChan)
	go sellTicket("售票口2", ticketChan)
	go sellTicket("售票口3", ticketChan)
	go sellTicket("售票口4", ticketChan)

	for i := ticket; i > 0; i-- {
		ticketChan <- i
	}

	time.Sleep(10 * time.Second)

}

func sellTicket(name string, ticketChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for {
		// 拿票
		ticket := <-ticketChan
		if ticket > 0 {
			// 睡眠，增大错误发生的概率
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// 售票
			fmt.Println(name, "售出：", ticket)
			// 减票
			ticket--
		} else {
			fmt.Println(name, "售罄，没票了...")
			break
		}
	}
}
