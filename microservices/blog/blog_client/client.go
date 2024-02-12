package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//insertBlog(c)
	readBlog(c)
}

func readBlog(c blogpb.BlogServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")
	req := &blogpb.ReadBlogRequest{
		BlogId: "65c907a3da60194db6d6fbc2",
	}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {

		co := status.Code(err)

		if co == codes.NotFound {
			fmt.Printf("Blog not found")
			return
		}

		fmt.Printf("Error while calling ReadBlog RPC: %v", err)
		return
	}

	fmt.Printf("The blog is %v", res)
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
