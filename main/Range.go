package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// 数组的range
	for i, v := range pow {
		println("2**", i, "=", v)
	}

	// 字符串的range
	for i, c := range "abcdefg" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}

	// map的range
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
	// read key only
	for k := range m {
		fmt.Printf("key: %s\n", k)
	}
	// read value only
	for _, v := range m {
		fmt.Printf("value: %d\n", v)
	}

	// channel的range
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)

	for i := range c {
		fmt.Println(i)
	}

}
