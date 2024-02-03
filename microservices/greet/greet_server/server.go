package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/greet/greetpb"
	"log"
	"net"
	"time"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	time := time.Now()
	fmt.Printf("Greet function was invoked with %v: %v\n", req, time)
	req.GetGreeting()
	firsName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firsName + "-" + lastName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	// create listener on port 50051

	fmt.Printf("Hello, World!\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
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
