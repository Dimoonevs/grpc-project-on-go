package main

import (
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpServe := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpServe, &helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpServe.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
