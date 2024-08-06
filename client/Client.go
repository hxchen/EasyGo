package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var replay string
	err = client.Call("LoginService.Login", "request", &replay)
	if err != nil {
		log.Fatal("Login error:", err)
	}
	log.Println("LoginService.Login:", replay)
}
