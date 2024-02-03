package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/greet/greetpb"
	"log"
)

func main() {

	fmt.Printf("I'm a client\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	fmt.Printf("Created client: %f", c)

}
