package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/calculator/calculatorpb"
	"io"
	"log"
	"net"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC\n")

	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	result := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func (s *server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	fmt.Printf("Received PrimeNumberDecomposition RPC\n, %v", req.GetNumber())

	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		tmp := number % divisor
		if tmp == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			}
			stream.Send(res)
			number = number / divisor
		} else {
			divisor = divisor + 1
		}
	}
	return nil
}

func (s *server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC\n")

	sum := int32(0)
	count := int32(0)
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				average := float64(sum) / float64(count)
				return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
					Average: int64(average),
				})
			}
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}

	return nil
}

func main() {
	fmt.Printf("Calculator Server\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server  %v", err)
	}
}
