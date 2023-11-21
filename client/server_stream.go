package main

import (
	"context"
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Streaming error: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Streaming error: %v", err)
		}
		log.Printf("Received: %s", res.Message)
	}
	log.Println("Streaming finished")
}
