package main

import (
	"io"
	"log"

	pb "github.com/Prototype-1/grpc-sample1/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	log.Println("Bidirectional streaming started...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed the stream.")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message from client: %v\n", err)
			return err
		}

		log.Printf("Received message from client: %s\n", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			log.Printf("Error sending message to client: %v\n", err)
			return err
		}
		log.Printf("Sent response to client: %s\n", res.Message)
	}
}
