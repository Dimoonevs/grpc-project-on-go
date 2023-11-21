package main

import (
	"context"
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, close := context.WithTimeout(context.Background(), time.Second)
	defer close()

	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
