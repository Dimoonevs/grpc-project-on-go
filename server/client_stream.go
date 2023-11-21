package main

import (
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var massages []string
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: massages})
		}
		if err != nil {
			return err
		}
		log.Printf("Get request with name: %s", res.Name)
		massages = append(massages, "Hello ", res.Name)
	}
}
