package main

import (
	pb "github.com/youtangai/HelloServerStreaming/proto"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
)

type helloServerStreamingService struct {}

func (hello *helloServerStreamingService) Greet(req *pb.HelloRequest, stream pb.HelloServerStreamingService_GreetServer) error {
	names := []string{"nagai", "nanaumi", "hiroto","taichi", "miyoshi"}
	for _, name := range names {
		if err := stream.Send(&pb.HelloResponse{Message: name}); err != nil {
			return err
		}
		log.Printf("send %s to client!\n", name)
	}
	return nil
}

func main() {
	fmt.Println("start streaming server!!")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServerStreamingServiceServer(grpcServer, &helloServerStreamingService{})
	grpcServer.Serve(lis)
}
