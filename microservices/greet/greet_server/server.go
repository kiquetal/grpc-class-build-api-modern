package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/greet/greetpb"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {

	// create listener on port 50051

	fmt.Printf("Hello, World!\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	lis.Accept()
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
