package main

import (
	"context"
	"io"
	"log"

	pb "github.com/xudong7/grpc-test/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("server streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("failed to call SayHelloServerStream: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive message from server: %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished")
}
