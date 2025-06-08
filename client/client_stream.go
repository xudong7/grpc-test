package main

import (
	"context"
	"log"
	"time"

	pb "github.com/xudong7/grpc-test/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to call SayHelloClientStreaming: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send message to server: %v", err)
		}
		log.Printf("sent name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("failed to receive response from server: %v", err)
	}
	log.Printf("Response from server: %v", res.Messages)
}
