package main

import (
	"context"
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	strem, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("cloud not send names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRquest{Name: name}
		if err := strem.Send(req); err != nil {
			log.Fatalf("Error while send %v", err)
		}
		log.Printf("Send the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := strem.CloseAndRecv()
	log.Println("Streaming finished")
	if err != nil {
		log.Fatalf("Error while reciving %v", err)
	}
	log.Printf("%v", res.Messages)
}
