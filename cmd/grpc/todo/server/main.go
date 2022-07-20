package main

import (
	protobuf "api-research/generated/grpc/todo/v1"
	"api-research/pkg/interfaces/grpc/todo/v1/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = 50051

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protobuf.RegisterTodoManagerServer(grpcServer, server.NewTodoServer())
	log.Printf("server listening at %v", listener.Addr())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
