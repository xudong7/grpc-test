package main

import (
	"io"
	"log"

	pb "github.com/xudong7/grpc-test/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil // End of stream, no more messages to process
		}
		if err != nil {
			return err // Handle any other errors
		}
		log.Printf("got request with name: %v", req.Name)
		// Create a response message
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err // Handle send error
		}
	}
}
