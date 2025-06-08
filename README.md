# An example using grpc and go to demonstrate client-server communication

## Overview

This repository contains an example of a gRPC client-server application written in Go. The server exposes a simple service that the client can call to demonstrate communication between the two.

## Prerequisites

- Go installed on your machine (version 1.16 or later)
- `protoc` compiler installed for generating gRPC code
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins installed for Go
  - You can install them using the following commands:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- Edit `~/.bashrc` or `~/.zshrc` to add the Go binaries to your PATH:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

```bash
source ~/.bashrc
# or
source ~/.zshrc
```

now you can check if the plugins are installed correctly:

```bash
protoc --version
# should output something like "libprotoc 3.x.x"
```

## Run

1. Clone the repository:

```bash
git clone https://github.com/xudong7/grpc-test.git
cd grpc-test
```

2. Generate the gRPC code from the `.proto` file:

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

3. Build the server and client:

```bash
cd server
go run *.go
```

In another terminal, build and run the client:

```bash
cd client
go run *.go
```

or you can use `Makefile` to generate proto code and run both server and client:

```bash
make protoc
```

```bash
make srun
```

In another terminal, run the client:

```bash
make crun
```
