package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Prototype-1/grpc-sample1/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming has started...")
stream, err := 	client.SayHelloBidirectionalStreaming(context.Background())
if err != nil {
	log.Fatalf("Could not send names: %v", err)
}

waitCh := make(chan struct{})
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
	close(waitCh)
}()

for _, name := range names.Names {
	req := &pb.HelloRequest{
		Name: name,
	}
	if err := stream.Send(req); err != nil {
		log.Fatalf("Error while sending %v", err)
	}
	time.Sleep((2 * time.Second))
}
stream.CloseSend()
<-waitCh
log.Printf("Bidirectional streaming finished...")
}