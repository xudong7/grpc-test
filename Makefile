# Makefile for gRPC example in Go

protoc: # used to generate Go code from .proto files
# if not successful, install protoc-gen-go and protoc-gen-go-grpc
# maybe you need to export GOPATH to your environment
# export PATH=$PATH:$(go env GOPATH)/bin
	protoc --go_out=. --go-grpc_out=. proto/greet.proto

srun: # run server
	go run server/*.go

crun: # run client
	go run client/*.go	

help:
	@echo "Makefile commands:"
	@echo "  protoc - Generate Go code from .proto files"
	@echo "  srun   - Run the server"
	@echo "  crun   - Run the client"
	@echo "  help   - Show this help message"

.PHONY: protoc srun crun help