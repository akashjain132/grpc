package main

import (
	"context"
	"fmt"
	"grpc/amazon"
	"log"
	"net"

	"google.golang.org/grpc"
)

// server to implement sum service
type server struct {
	amazon.UnimplementedProductServiceServer
}

func (*server) GetProduct(ctx context.Context, req *amazon.ProductID) (*amazon.Product, error) {

	return &amazon.Product{Id: "1234", Name: "MacPro", Description: "Mac book", Weight: 1000, Price: 100000}, nil
}

func main() {
	fmt.Println("Starting server")

	// Listening to tcp network on address localhost with port 60001
	lis, err := net.Listen("tcp", "0.0.0.0:60001")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// create new grpc service
	s := grpc.NewServer()
	// register sum service with grpc
	amazon.RegisterProductServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to listen to the server %v", err)
	}

	log.Println("Connected to the server")
}
