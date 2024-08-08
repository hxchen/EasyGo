package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
	name string
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Printf("I am %s, I can call you!\n", nokiaPhone.name)
}

type IPhone struct {
	name string
}

func (iPhone IPhone) call() {
	fmt.Printf("I am %s, I can call you!\n", iPhone.name)
}

func main() {
	var phone Phone

	phone = NokiaPhone{name: "Nokia"}
	phone.call()

	phone = IPhone{name: "iPhone"}
	phone.call()
}
