package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpc/learning/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

var collection *mongo.Collection

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("Blog Service Started\n")

	var options = options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:      "mydbadmin",
		Password:      "admin",
		AuthSource:    "mydb",
		AuthMechanism: "SCRAM-SHA-256",
	})
	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging mongo client: %v", err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database("mydb").Collection("blog")

	if collection == nil {
		log.Fatalf("Error getting collection: %v", err)
	}

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server %v", err)
		}

	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("Stopping the server")
	lis.Close()
	fmt.Println("Closing the listener")
	s.Stop()
	fmt.Println("Stopping the server")
}
