package main

import (
	pb "github.com/Dimoonevs/gRPC-test-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Dima", "Ostap", "Petr"},
	}

	//callSayHello(client)

	//callSayHelloServerStream(client, names)

	//callSayHelloClientStream(client, names)

	callSayHelloBidirectionalStream(client, names)
}
