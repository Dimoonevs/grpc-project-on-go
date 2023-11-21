package main

import (
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBiDirectional(stream pb.GreetService_SayHelloBiDirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		resp := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}

	}
}
