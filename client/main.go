package main

import (
	"log"

	pb "github.com/xudong7/grpc-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost"
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial(addr+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{Names: []string{
		"dunjia",
		"bob",
		"alice",
	}}

	callSayHello(client)
	callSayHelloServerStreaming(client, names)
	callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
