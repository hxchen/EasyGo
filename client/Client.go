package main

import (
	"context"
	rpc "github.com/hxchen/EasyGo/api/protobuf-spec"
	"google.golang.org/grpc"
	"log"
)

const PORT = ":1234"

func main() {
	conn, err := grpc.Dial(PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("get an error : %v\n", err)
	}
	defer conn.Close()

	client := rpc.NewHelloServiceClient(conn)

	resp, err := client.Say(context.Background(), &rpc.HelloRequest{
		Name: "this is client",
	})
	if err != nil {
		log.Fatalf("invoke error \n")
	}

	log.Printf("resp : %s\n", resp.GetMessage())
}
