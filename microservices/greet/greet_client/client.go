package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/learning/greet/greetpb"
	"io"
	"log"
)

func main() {

	fmt.Printf("I'm a client\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	//	doUnary(c)
	doStream(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do a Unary RPC...\n")

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sunholiday",
			LastName:  "Sun",
		},
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doStream(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do a Server Streaming RPC...\n")
	res, err := c.ManyTimesGreet(context.Background(), &greetpb.ManyGreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sunholiday",
			LastName:  "Sun",
		}})
	if err != nil {
		log.Fatalf("Error while calling ManyTimesGreet RPC: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
			break
		}
		log.Printf("Response from ManyTimesGreet: %v", msg.GetResult())
	}
	log.Printf("Finished to do a Server Streaming RPC...\n")
}
