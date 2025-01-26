package main

import (
	"log"
pb "github.com/Prototype-1/grpc-sample1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
conn, err := grpc.NewClient("localhost" + port, grpc.WithTransportCredentials(insecure.NewCredentials()))

if err != nil {
	log.Fatalf("Did not connect: %v", err)
}
defer conn.Close()

client := pb.NewGreetServiceClient(conn)

names := &pb.NamesList{
	Names: [] string {"Aswin", "Anju", "Mebin"},
}

callSayHello(client)

callSayHelloServerStream(client, names)
}