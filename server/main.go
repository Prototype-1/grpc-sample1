// package main

// import (
// 	"log"
// 	"net"
// 	pb "github.com/Prototype-1/grpc-sample1/proto"
// 	"google.golang.org/grpc"
// )

// const (
// 	port = ":8080"
// )

// type helloServer struct {
// 	pb.GreetServiceServer
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("Failed to start the server %v", err)
// 	}
// 	grpcServer := grpc.NewServer()
// 	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
// 	log.Printf("Server started at %v", lis.Addr())
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to start: %v", err)
// 	}
// }

package main

import (
	"log"
	"net"

	pb "github.com/Prototype-1/grpc-sample1/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.UnimplementedGreetServiceServer // Avoids breaking due to future protobuf changes
}

func main() {
	// Listen on the specified port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the GreetService server implementation
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Server is running on %v\n", port)

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
