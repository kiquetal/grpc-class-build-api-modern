package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc/learning/blog/blogpb"
	"io"
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
	/*	insertBlog(c, &blogpb.Blog{
			AuthorId: "JR.Tolkien",
			Title:    "Lord of the Rings",
			Content:  "Here's Johnny!"})
		insertBlog(c,
			&blogpb.Blog{
				AuthorId: "Stephen King2",
				Title:    "The Shining",
				Content:  "Here's Johnny!"})
	*/
	//readBlog(c)
	//updateBlog(c)
	//deleteBlog(c)
	listBlog(c)
}

func listBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blogpb.ListBlogRequest{}
	res, err := c.ListBlog(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while calling ListBlog RPC: %v", err)
		return
	}
	for {
		msg, err := res.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Error while reading stream: %v", err)
			break
		}
		fmt.Println("Blog found: %v", msg)
	}

}

func deleteBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blogpb.ReadBlogRequest{
		BlogId: "65c907a3da60194db6d6fbc2",
	}
	res, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		co := status.Code(err)

		if co == codes.NotFound {
			fmt.Printf("Blog not found")
			return
		}

		fmt.Printf("Error while calling DeleteBlog RPC: %v", err)
		return

	}
	fmt.Printf("Deleted blog with Id: %v", res)

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

func insertBlog(c blogpb.BlogServiceClient, blog *blogpb.Blog) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &blogpb.CreateBlogRequest{
		Blog: blog,
	}

	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		fmt.Printf("Error while calling CreateBlog RPC: %v", err)
		return
	}

	fmt.Printf("The created id blog is %v", res.GetBlog().GetId())
}
