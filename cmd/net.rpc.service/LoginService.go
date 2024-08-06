package main

import (
	"log"
	"net"
	"net/rpc"
)

type LoginService struct{}

func (ls *LoginService) Login(request string, reply *string) error {
	*reply = "LoginService.Login"
	return nil
}

func main() {
	rpc.RegisterName("LoginService", new(LoginService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}
