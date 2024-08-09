package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func a() {
	for i := 1; i < 10; i++ {
		wg.Done()
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		wg.Done()
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(18)
	go a()
	go b()
	wg.Wait()
}
