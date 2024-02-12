package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/learning/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	requestBlog := req.GetBlog()
	if requestBlog == nil {
		return nil, fmt.Errorf("Request Blog is empty")
	}
	data := blogItem{
		AuthorID: requestBlog.GetAuthorId(),
		Content:  requestBlog.GetContent(),
		Title:    requestBlog.GetTitle(),
	}

	// check if the data already exists

	var filter = bson.M{"author_id": data.AuthorID}
	fmt.Println("Filter: ", filter)
	var data2 blogItem
	resultFind := collection.FindOne(context.Background(), filter).Decode(&data2)

	if resultFind != nil {
		if errors.Is(resultFind, mongo.ErrNoDocuments) {
			fmt.Println("No Data Found")
		} else {
			return nil, fmt.Errorf("Error finding blog: %v", resultFind)

		}

	}
	if data2 != (blogItem{}) {
		return nil, errors.New("Blog already exists")

	}
	// insert the data
	result, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, fmt.Errorf("Error inserting blog: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("Error converting to OID")
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.String(),
			AuthorId: data.AuthorID,
			Content:  data.Content,
			Title:    data.Title,
		},
	}, nil

}

func (s *server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.Blog, error) {
	blogID := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error converting to OID: %v", err)
	}
	filter := bson.M{"_id": oid}
	var data blogItem
	result := collection.FindOne(context.Background(), filter).Decode(&data)
	if result != nil {
		if errors.Is(result, mongo.ErrNoDocuments) {
			return nil, status.Errorf(codes.NotFound, "Blog with ID %v not found", blogID)
		} else {
			return nil, fmt.Errorf("error finding blog: %v", result)
		}
	}
	return &blogpb.Blog{
		Id:       data.ID,
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}, nil
}

var collection *mongo.Collection

type blogItem struct {
	ID       string `bson:"_id,omitempty"`
	AuthorID string `bson:"author_id"`
	Content  string `bson:"content"`
	Title    string `bson:"title"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("Blog Service Started\n")

	var optionsO = options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:      "mydbadmin",
		Password:      "admin",
		AuthSource:    "mydb",
		AuthMechanism: "SCRAM-SHA-256",
	})
	client, err := mongo.Connect(context.Background(), optionsO)
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

	var opts []grpc.ServerOption
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
	err = lis.Close()
	if err != nil {
		return
	}
	fmt.Println("Closing the listener")
	s.Stop()
	fmt.Println("Client Disconnected")
	err = client.Disconnect(context.Background())
	if err != nil {
		return
	}
	fmt.Println("MongoDB Disconnected")
}
