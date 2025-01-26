package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Prototype-1/grpc-sample1/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for {
	msg, err := stream.Recv()
	if err ==io.EOF {
		break
	}
	if err != nil {
		log.Fatalf("Error while streaming: %v", err)
	}
	log.Println(msg)
	}
	log.Printf("Stream finished")
}

