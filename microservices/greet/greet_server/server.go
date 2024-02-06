package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/greet/greetpb"
	"io"
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

func (s *server) ManyTimesGreet(req *greetpb.ManyGreetRequest, stream greetpb.GreetService_ManyTimesGreetServer) error {
	fmt.Printf("ManyTimesGreet function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + fmt.Sprint(i) + " " + lastName
		res := &greetpb.ManyGreetResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {

	fmt.Printf("LongGreet function was invoked with a streaming request\n")

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.GreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!" + "\n"
	}
	return nil
}

func (s *server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreatEveryone function was invoked with a streaming request\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v",
				err)
			return err
		}
		firstName := req.GetGreeting().FirstName
		lastName := req.GetGreeting().LastName
		result := "Hello " + firstName + " " + lastName + "!"
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
		}
	}
	return nil
}
func main() {

	// create listener on port 50051

	fmt.Printf("Greet Server\n")
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
