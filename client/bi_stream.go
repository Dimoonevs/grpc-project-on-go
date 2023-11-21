package main

import (
	"context"
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Bidirectional Streaming Started")
	stream, err := client.SayHelloBiDirectional(context.Background())
	if err != nil {
		log.Fatalf("cloud not send names: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRquest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while send %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("Bidirectional Streaming End")
}
