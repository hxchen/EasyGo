package main

import (
	"context"
	hi "github.com/hxchen/EasyGo/api/protobuf-spec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":1234"
)

type server struct {
	hi.HelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *hi.HelloRequest) (*hi.HelloResponse, error) {
	log.Printf("Received: %s", in.Name)
	return &hi.HelloResponse{Message: "Hello " + in.Name}, nil
}

func listenAndService() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen, %v", err)
	}
	s := grpc.NewServer()
	hi.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	log.Print("the rpc server is started up\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

func main() {
	listenAndService()
}
