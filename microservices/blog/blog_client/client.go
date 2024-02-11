package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/learning/blog/blogpb"
)

func main() {
	fmt.Printf("Client for the blog service\n")

	//obtain the conneciton from grpc
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		fmt.Printf("Error connecting to the server: %v", err)
	}
	defer client.Close()

	c := blogpb.NewBlogServiceClient(client)
	insertBlog(c)
}

func insertBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Stephen King",
			Title:    "The Shining",
			Content:  "Here's Johnny!",
		},
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while calling CreateBlog RPC: %v", err)
		return
	}

	fmt.Printf("The created id blog is %v", res.GetBlog().GetId())
}
