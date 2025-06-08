package main

import (
	"context"
	"log"
	"time"

	pb "github.com/xudong7/grpc-test/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response from server: %s", res.Message)
}
