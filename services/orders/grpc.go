package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis,err := net.Listen("tcp",s.addr)
	if err !=nil {
		log.Fatal("failed to listen: %v",err)
	}
	gRPCServer := grpc.NewServer()
	// register our grpc services
	log.Println("Starting gRPC server on",s.addr)
	return gRPCServer.Serve(lis)
}