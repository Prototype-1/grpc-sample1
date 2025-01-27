// package main

// import (
// 	"io"
// 	"log"

// 	pb "github.com/Prototype-1/grpc-sample1/proto"
// )

// func (s *helloServer) SayHelloBidirectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		log.Printf("Got request with name: %v", req.Name)
// 		res := &pb.HelloResponse{
// 			Message: "Hello " + req.Name,
// 		}
// 		if err := stream.Send(res); err != nil {
// 			return err
// 		}
// 	}
// }

package main

import (
	"io"
	"log"

	pb "github.com/Prototype-1/grpc-sample1/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	log.Println("Bidirectional streaming started...")
	for {
		// Receive a message from the client
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed the stream.")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message from client: %v\n", err)
			return err
		}

		// Log the received message
		log.Printf("Received message from client: %s\n", req.Name)

		// Create and send a response
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
