package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/learning/calculator/calculatorpb"
	"io"
	"log"
	"math"
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

func (s *server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {

	fmt.Printf("Received FindMaximum RPC\n")
	maximum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		reqNumber := req.GetNumber()
		if reqNumber > maximum {
			maximum = reqNumber
			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: maximum,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
		}
	}

	return nil
}

func (s *server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("Received SquareRoot RPC\n")
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received bad argument %v", number),
		)
	}

	// calculate square root

	return &calculatorpb.SquareRootResponse{
		NumberRoot: float32(math.Sqrt(float64(number))),
	}, nil

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
