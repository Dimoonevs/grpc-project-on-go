package main

import (
	"context"
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
)

func (s *helloServer) SayHello(ctx context.Context, input *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello"}, nil
}
