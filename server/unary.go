package main

import (
	"context"

	pb "github.com/xudong7/grpc-test/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello, World!",
	}, nil
}
