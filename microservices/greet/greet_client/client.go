package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/learning/greet/greetpb"
	"io"
	"log"
	"time"
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
	//	doStream(c)
	//	doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do a LongGreet RPC...\n")
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday",
				LastName:  "Sun",
			}},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday2",
				LastName:  "Sun2",
			}},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday3",
				LastName:  "Sun3",
			}},
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Errorf("errror %v", err)

	}
	fmt.Printf("LongGreet Response: \n%v\n", res.Result)

}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Printf("Starting to do a BiDi Streaming RPC...\n")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}
	// Send a bunch of messages to the client (go routine)

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday",
				LastName:  "Sun",
			}},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday2",
				LastName:  "Sun2",
			}},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday3",
				LastName:  "Sun3",
			}},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sunholiday4",
				LastName:  "Sun4",
			}},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)

		}
		stream.CloseSend()

	}()

	go func() {
		for {
			stream, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				close(waitc)
				break
			}
			fmt.Printf("Received: %v\n", stream.GetResult())

		}
	}()

	<-waitc
}
