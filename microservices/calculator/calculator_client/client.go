package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/learning/calculator/calculatorpb"
	"io"
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
	doClientStream(c)
	//	doBiDiStreaming(c)

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
