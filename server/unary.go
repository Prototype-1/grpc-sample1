package main

import (
	"context"
	pb "github.com/Prototype-1/grpc-sample1/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse {
		Message: "Hello",
	}, nil
}