package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/learning/calculator/calculatorpb"
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
	doUnary(c)

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
