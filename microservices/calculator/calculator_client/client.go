package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/learning/calculator/calculatorpb"
	"io"
	"time"
)

func main() {

	fmt.Printf("Calculator Client\n")

	cc, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		fmt.Printf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//	doUnary(c)
	//	doStreamDecomposition(c)
	//doClientStream(c)
	doBiDiStreaming(c)

}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a BiDi Streaming RPC...\n")
	// we create a stream by invoking the client
	requests := []*calculatorpb.FindMaximumRequest{
		&calculatorpb.FindMaximumRequest{
			Number: 4,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 3,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 5,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 1,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 6,
		},
	}
	waitc := make(chan struct{})
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		fmt.Printf("Error while creating stream: %v", err)
		return
	}
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			// read
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetMaximum())

		}
		close(waitc)

	}()

	<-waitc
}

func doClientStream(c calculatorpb.CalculatorServiceClient) {

	fmt.Printf("Starting to do a Client Streaming RPC...\n")
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		fmt.Printf("Error while calling")

	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		fmt.Printf("Sending number: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("Error while receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("The Average is: %v\n", res.GetAverage())

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a Unary RPC...\n")

	res, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		FirstNumber:  23,
		SecondNumber: 10,
	})

	if err != nil {
		fmt.Printf("Error while calling Greet RPC: %v", err)
	}
	fmt.Printf("Response from Calculator: %v", res.SumResult)

}

func doStreamDecomposition(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a PrimeNumberDecomposition RPC...\n")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 155455664,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			fmt.Printf("Error while reading stream: %v", err)
		}
		fmt.Printf("Response from PrimeNumberDecomposition: %v\n", msg.GetPrimeFactor())
	}
}
