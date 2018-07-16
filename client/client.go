package main

import (
	pb "github.com/youtangai/HelloServerStreaming/proto"
	"golang.org/x/net/context"
	"time"
	"log"
	"io"
	"fmt"
	"google.golang.org/grpc"
)

func runGreet(client pb.HelloServerStreamingServiceClient) {
	req := &pb.HelloRequest{Message: "hello"}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	stream, err := client.Greet(ctx, req)

	if err != nil {
		log.Fatalf("cennot exec greet: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot recieve: %v", err)
		}
		fmt.Println(resp.Message)
	}
}

func main() {
	fmt.Println("start client!!")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect grpc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServerStreamingServiceClient(conn)

	runGreet(client)
}
